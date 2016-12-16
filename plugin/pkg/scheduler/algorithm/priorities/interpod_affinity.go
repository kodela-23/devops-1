/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package priorities

import (
	"sync"

	"github.com/golang/glog"
<<<<<<< HEAD
	"k8s.io/kubernetes/pkg/api"
=======
	"k8s.io/kubernetes/pkg/api/v1"
	metav1 "k8s.io/kubernetes/pkg/apis/meta/v1"
>>>>>>> 55b413f... Do not create selector and namespaces in a loop where possible
	"k8s.io/kubernetes/pkg/util/workqueue"
	"k8s.io/kubernetes/plugin/pkg/scheduler/algorithm"
	"k8s.io/kubernetes/plugin/pkg/scheduler/algorithm/predicates"
	priorityutil "k8s.io/kubernetes/plugin/pkg/scheduler/algorithm/priorities/util"
	schedulerapi "k8s.io/kubernetes/plugin/pkg/scheduler/api"
	"k8s.io/kubernetes/plugin/pkg/scheduler/schedulercache"
)

type InterPodAffinity struct {
	info                  predicates.NodeInfo
	nodeLister            algorithm.NodeLister
	podLister             algorithm.PodLister
	hardPodAffinityWeight int
	failureDomains        priorityutil.Topologies
}

func NewInterPodAffinityPriority(
	info predicates.NodeInfo,
	nodeLister algorithm.NodeLister,
	podLister algorithm.PodLister,
	hardPodAffinityWeight int,
	failureDomains []string) algorithm.PriorityFunction {
	interPodAffinity := &InterPodAffinity{
		info:                  info,
		nodeLister:            nodeLister,
		podLister:             podLister,
		hardPodAffinityWeight: hardPodAffinityWeight,
		failureDomains:        priorityutil.Topologies{DefaultKeys: failureDomains},
	}
	return interPodAffinity.CalculateInterPodAffinityPriority
}

type podAffinityPriorityMap struct {
	sync.Mutex

	// nodes contain all nodes that should be considered
	nodes []*api.Node
	// counts store the mapping from node name to so-far computed score of
	// the node.
	counts map[string]float64
	// failureDomains contain default failure domains keys
	failureDomains priorityutil.Topologies
	// The first error that we faced.
	firstError error
}

func newPodAffinityPriorityMap(nodes []*api.Node, failureDomains priorityutil.Topologies) *podAffinityPriorityMap {
	return &podAffinityPriorityMap{
		nodes:          nodes,
		counts:         make(map[string]float64, len(nodes)),
		failureDomains: failureDomains,
	}
}

func (p *podAffinityPriorityMap) setError(err error) {
	p.Lock()
	defer p.Unlock()
	if p.firstError == nil {
		p.firstError = err
	}
}

func (p *podAffinityPriorityMap) processTerm(term *api.PodAffinityTerm, podDefiningAffinityTerm, podToCheck *api.Pod, fixedNode *api.Node, weight float64) {
	namespaces := priorityutil.GetNamespacesFromPodAffinityTerm(podDefiningAffinityTerm, term)
	selector, err := metav1.LabelSelectorAsSelector(term.LabelSelector)
	if err != nil {
		p.setError(err)
		return
	}
	match := priorityutil.PodMatchesTermsNamespaceAndSelector(podToCheck, namespaces, selector)
	if match {
		func() {
			p.Lock()
			defer p.Unlock()
			for _, node := range p.nodes {
				if p.failureDomains.NodesHaveSameTopologyKey(node, fixedNode, term.TopologyKey) {
					p.counts[node.Name] += weight
				}
			}
		}()
	}
}

func (p *podAffinityPriorityMap) processTerms(terms []api.WeightedPodAffinityTerm, podDefiningAffinityTerm, podToCheck *api.Pod, fixedNode *api.Node, multiplier int) {
	for i := range terms {
		term := &terms[i]
		p.processTerm(&term.PodAffinityTerm, podDefiningAffinityTerm, podToCheck, fixedNode, float64(term.Weight*int32(multiplier)))
	}
}

// compute a sum by iterating through the elements of weightedPodAffinityTerm and adding
// "weight" to the sum if the corresponding PodAffinityTerm is satisfied for
// that node; the node(s) with the highest sum are the most preferred.
// Symmetry need to be considered for preferredDuringSchedulingIgnoredDuringExecution from podAffinity & podAntiAffinity,
// symmetry need to be considered for hard requirements from podAffinity
func (ipa *InterPodAffinity) CalculateInterPodAffinityPriority(pod *api.Pod, nodeNameToInfo map[string]*schedulercache.NodeInfo, nodes []*api.Node) (schedulerapi.HostPriorityList, error) {
	affinity, err := api.GetAffinityFromPodAnnotations(pod.Annotations)
	if err != nil {
		return nil, err
	}
	hasAffinityConstraints := affinity != nil && affinity.PodAffinity != nil
	hasAntiAffinityConstraints := affinity != nil && affinity.PodAntiAffinity != nil

	allNodeNames := make([]string, 0, len(nodeNameToInfo))
	for name := range nodeNameToInfo {
		allNodeNames = append(allNodeNames, name)
	}

	// convert the topology key based weights to the node name based weights
	var maxCount float64
	var minCount float64
	// priorityMap stores the mapping from node name to so-far computed score of
	// the node.
	pm := newPodAffinityPriorityMap(nodes, ipa.failureDomains)

	processPod := func(existingPod *api.Pod) error {
		existingPodNode, err := ipa.info.GetNodeInfo(existingPod.Spec.NodeName)
		if err != nil {
			return err
		}
		existingPodAffinity, err := api.GetAffinityFromPodAnnotations(existingPod.Annotations)
		if err != nil {
			return err
		}
		existingHasAffinityConstraints := existingPodAffinity != nil && existingPodAffinity.PodAffinity != nil
		existingHasAntiAffinityConstraints := existingPodAffinity != nil && existingPodAffinity.PodAntiAffinity != nil

		if hasAffinityConstraints {
			// For every soft pod affinity term of <pod>, if <existingPod> matches the term,
			// increment <pm.counts> for every node in the cluster with the same <term.TopologyKey>
			// value as that of <existingPods>`s node by the term`s weight.
			terms := affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution
			pm.processTerms(terms, pod, existingPod, existingPodNode, 1)
		}
		if hasAntiAffinityConstraints {
			// For every soft pod anti-affinity term of <pod>, if <existingPod> matches the term,
			// decrement <pm.counts> for every node in the cluster with the same <term.TopologyKey>
			// value as that of <existingPod>`s node by the term`s weight.
			terms := affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution
			pm.processTerms(terms, pod, existingPod, existingPodNode, -1)
		}

		if existingHasAffinityConstraints {
			// For every hard pod affinity term of <existingPod>, if <pod> matches the term,
			// increment <pm.counts> for every node in the cluster with the same <term.TopologyKey>
			// value as that of <existingPod>'s node by the constant <ipa.hardPodAffinityWeight>
			if ipa.hardPodAffinityWeight > 0 {
				terms := existingPodAffinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution
				// TODO: Uncomment this block when implement RequiredDuringSchedulingRequiredDuringExecution.
				//if len(existingPodAffinity.PodAffinity.RequiredDuringSchedulingRequiredDuringExecution) != 0 {
				//	terms = append(terms, existingPodAffinity.PodAffinity.RequiredDuringSchedulingRequiredDuringExecution...)
				//}
				for _, term := range terms {
					pm.processTerm(&term, existingPod, pod, existingPodNode, float64(ipa.hardPodAffinityWeight))
				}
			}
			// For every soft pod affinity term of <existingPod>, if <pod> matches the term,
			// increment <pm.counts> for every node in the cluster with the same <term.TopologyKey>
			// value as that of <existingPod>'s node by the term's weight.
			terms := existingPodAffinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution
			pm.processTerms(terms, existingPod, pod, existingPodNode, 1)
		}
		if existingHasAntiAffinityConstraints {
			// For every soft pod anti-affinity term of <existingPod>, if <pod> matches the term,
			// decrement <pm.counts> for every node in the cluster with the same <term.TopologyKey>
			// value as that of <existingPod>'s node by the term's weight.
			terms := existingPodAffinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution
			pm.processTerms(terms, existingPod, pod, existingPodNode, -1)
		}
		return nil
	}
	processNode := func(i int) {
		nodeInfo := nodeNameToInfo[allNodeNames[i]]
		if hasAffinityConstraints || hasAntiAffinityConstraints {
			// We need to process all the nodes.
			for _, existingPod := range nodeInfo.Pods() {
				if err := processPod(existingPod); err != nil {
					pm.setError(err)
				}
			}
		} else {
			// The pod doesn't have any constraints - we need to check only existing
			// ones that have some.
			for _, existingPod := range nodeInfo.PodsWithAffinity() {
				if err := processPod(existingPod); err != nil {
					pm.setError(err)
				}
			}
		}
	}
	workqueue.Parallelize(16, len(allNodeNames), processNode)
	if pm.firstError != nil {
		return nil, pm.firstError
	}

	for _, node := range nodes {
		if pm.counts[node.Name] > maxCount {
			maxCount = pm.counts[node.Name]
		}
		if pm.counts[node.Name] < minCount {
			minCount = pm.counts[node.Name]
		}
	}

	// calculate final priority score for each node
	result := make(schedulerapi.HostPriorityList, 0, len(nodes))
	for _, node := range nodes {
		fScore := float64(0)
		if (maxCount - minCount) > 0 {
			fScore = 10 * ((pm.counts[node.Name] - minCount) / (maxCount - minCount))
		}
		result = append(result, schedulerapi.HostPriority{Host: node.Name, Score: int(fScore)})
		if glog.V(10) {
			// We explicitly don't do glog.V(10).Infof() to avoid computing all the parameters if this is
			// not logged. There is visible performance gain from it.
			glog.V(10).Infof("%v -> %v: InterPodAffinityPriority, Score: (%d)", pod.Name, node.Name, int(fScore))
		}
	}
	return result, nil
}
