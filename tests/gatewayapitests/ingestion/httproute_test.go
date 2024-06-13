/*
 * Copyright 2023-2024 VMware, Inc.
 * All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package ingestion

import (
	"testing"

	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"

	akogatewayapilib "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/ako-gateway-api/lib"
	akogatewayapiobjects "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/ako-gateway-api/objects"
	akogatewayapitests "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/tests/gatewayapitests"
)

func TestHTTPRouteCUD(t *testing.T) {
	gatewayClassName := "gateway-class-01"
	gatewayName := "gateway-01"
	httpRouteName := "httproute-01"
	namespace := "default"
	ports := []int32{8080, 8081}
	key := "HTTPRoute" + "/" + namespace + "/" + httpRouteName
	gwkey := "Gateway/" + DEFAULT_NAMESPACE + "/" + gatewayName
	akogatewayapiobjects.GatewayApiLister().UpdateGatewayClass(gatewayClassName, true)

	akogatewayapitests.SetupGatewayClass(t, gatewayClassName, akogatewayapilib.GatewayController)
	t.Logf("Created GatewayClass %s", gatewayClassName)
	waitAndverify(t, "GatewayClass/gateway-class-01")

	listeners := akogatewayapitests.GetListenersV1(ports, false)
	akogatewayapitests.SetupGateway(t, gatewayName, namespace, gatewayClassName, nil, listeners)
	t.Logf("Created Gateway %s", gatewayName)
	waitAndverify(t, "Gateway/default/gateway-01")

	parentRefs := akogatewayapitests.GetParentReferencesV1([]string{gatewayName}, namespace, ports)
	hostnames := []gatewayv1.Hostname{"foo-8080.com", "foo-8081.com"}
	akogatewayapitests.SetupHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)
	waitAndverify(t, key)

	// update
	hostnames = []gatewayv1.Hostname{"foo-8080.com"}
	akogatewayapitests.UpdateHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)
	waitAndverify(t, key)

	// delete
	akogatewayapitests.TeardownHTTPRoute(t, httpRouteName, namespace)
	waitAndverify(t, key)
	akogatewayapitests.TeardownGateway(t, gatewayName, namespace)
	waitAndverify(t, gwkey)
}

func TestHTTPRouteHostnameInvalid(t *testing.T) {
	gatewayClassName := "gateway-class-02"
	gatewayName := "gateway-02"
	httpRouteName := "httproute-02"
	gwKey := "Gateway/" + DEFAULT_NAMESPACE + "/" + gatewayName
	gwClassKey := "GatewayClass/" + gatewayClassName
	namespace := "default"
	ports := []int32{8080}
	key := "HTTPRoute" + "/" + namespace + "/" + httpRouteName
	akogatewayapiobjects.GatewayApiLister().UpdateGatewayClass(gatewayClassName, true)

	akogatewayapitests.SetupGatewayClass(t, gatewayClassName, akogatewayapilib.GatewayController)
	t.Logf("Created GatewayClass %s", gatewayClassName)
	waitAndverify(t, gwClassKey)

	listeners := akogatewayapitests.GetListenersV1(ports, false)
	akogatewayapitests.SetupGateway(t, gatewayName, namespace, gatewayClassName, nil, listeners)
	t.Logf("Created Gateway %s", gatewayName)
	waitAndverify(t, gwKey)

	parentRefs := akogatewayapitests.GetParentReferencesV1([]string{gatewayName}, namespace, ports)
	hostnames := []gatewayv1.Hostname{"*.example.com"}
	akogatewayapitests.SetupHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)
	waitAndverify(t, "")

	// update
	hostnames = []gatewayv1.Hostname{"foo-8080.com"}
	akogatewayapitests.UpdateHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)
	waitAndverify(t, key)

	// delete
	akogatewayapitests.TeardownHTTPRoute(t, httpRouteName, namespace)
	waitAndverify(t, key)
	akogatewayapitests.TeardownGateway(t, gatewayName, DEFAULT_NAMESPACE)
	waitAndverify(t, gwKey)
	akogatewayapitests.TeardownGatewayClass(t, gatewayClassName)
	waitAndverify(t, gwClassKey)
}

func TestHTTPRouteGatewayNotPresent(t *testing.T) {
	gatewayClassName := "gateway-class-03"
	gatewayName := "gateway-03"
	httpRouteName := "httproute-03"
	gwKey := "Gateway/" + DEFAULT_NAMESPACE + "/" + gatewayName
	gwClassKey := "GatewayClass/" + gatewayClassName
	namespace := "default"
	ports := []int32{8080, 8081}
	key := "HTTPRoute" + "/" + namespace + "/" + httpRouteName
	akogatewayapiobjects.GatewayApiLister().UpdateGatewayClass(gatewayClassName, true)

	akogatewayapitests.SetupGatewayClass(t, gatewayClassName, akogatewayapilib.GatewayController)
	t.Logf("Created GatewayClass %s", gatewayClassName)
	waitAndverify(t, gwClassKey)

	parentRefs := akogatewayapitests.GetParentReferencesV1([]string{gatewayName}, namespace, ports)
	hostnames := []gatewayv1.Hostname{"foo-8080.com", "foo-8081.com"}
	akogatewayapitests.SetupHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)
	waitAndverify(t, "")

	// update
	listeners := akogatewayapitests.GetListenersV1(ports, false)
	akogatewayapitests.SetupGateway(t, gatewayName, namespace, gatewayClassName, nil, listeners)
	t.Logf("Created Gateway %s", gatewayName)
	waitAndverify(t, gwKey)
	hostnames = []gatewayv1.Hostname{"foo-8080.com"}
	akogatewayapitests.UpdateHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)
	waitAndverify(t, key)

	// delete
	akogatewayapitests.TeardownHTTPRoute(t, httpRouteName, namespace)
	waitAndverify(t, key)
	akogatewayapitests.TeardownGateway(t, gatewayName, DEFAULT_NAMESPACE)
	waitAndverify(t, gwKey)
	akogatewayapitests.TeardownGatewayClass(t, gatewayClassName)
	waitAndverify(t, gwClassKey)
}

func TestHTTPRouteWithBackendRefFilters(t *testing.T) {
	gatewayClassName := "gateway-class-hr-14"
	gatewayName := "gateway-hr-14"
	httpRouteName := "httproute-14"
	svcName := "avisvc-hr-057"
	namespace := "default"
	ports := []int32{8080}

	akogatewayapitests.SetupGatewayClass(t, gatewayClassName, akogatewayapilib.GatewayController)

	listeners := akogatewayapitests.GetListenersV1(ports)

	g := gomega.NewGomegaWithT(t)

	integrationtest.CreateSVC(t, DEFAULT_NAMESPACE, svcName, corev1.ProtocolTCP, corev1.ServiceTypeClusterIP, false)
	integrationtest.CreateEP(t, DEFAULT_NAMESPACE, svcName, false, false, "1.2.3")

	akogatewayapitests.SetupGateway(t, gatewayName, namespace, gatewayClassName, nil, listeners)
	g.Eventually(func() bool {
		gateway, err := akogatewayapitests.GatewayClient.GatewayV1().Gateways(namespace).Get(context.TODO(), gatewayName, metav1.GetOptions{})
		if err != nil || gateway == nil {
			t.Logf("Couldn't get the gateway, err: %+v", err)
			return false
		}
		return apimeta.FindStatusCondition(gateway.Status.Conditions, string(gatewayv1.GatewayConditionAccepted)) != nil
	}, 30*time.Second).Should(gomega.Equal(true))

	parentRefs := akogatewayapitests.GetParentReferencesV1([]string{gatewayName}, namespace, ports)
	hostnames := []gatewayv1.Hostname{"foo-8080.com", "foo-8081.com"}
	rule := akogatewayapitests.GetHTTPRouteRuleV1([]string{"/foo"}, []string{}, nil,
		[][]string{{svcName, DEFAULT_NAMESPACE, "8080", "1"}}, map[string][]string{"RequestHeaderModifier": {"add", "remove", "replace"}})
	rules := []gatewayv1.HTTPRouteRule{rule}

	akogatewayapitests.SetupHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, rules)
	g.Eventually(func() bool {
		httpRoute, err := akogatewayapitests.GatewayClient.GatewayV1().HTTPRoutes(namespace).Get(context.TODO(), httpRouteName, metav1.GetOptions{})
		if err != nil || httpRoute == nil {
			t.Logf("Couldn't get the HTTPRoute, err: %+v", err)
			return false
		}
		if len(httpRoute.Status.Parents) != len(ports) {
			return false
		}
		return apimeta.FindStatusCondition(httpRoute.Status.Parents[0].Conditions, string(gatewayv1.GatewayConditionAccepted)) != nil
	}, 30*time.Second).Should(gomega.Equal(true))

	conditionMap := make(map[string][]metav1.Condition)

	for _, port := range ports {
		conditions := make([]metav1.Condition, 0, 1)
		condition := metav1.Condition{
			Type:    string(gatewayv1.GatewayConditionAccepted),
			Reason:  string(gatewayv1.GatewayReasonAccepted),
			Status:  metav1.ConditionTrue,
			Message: "Parent reference is valid",
		}
		conditions = append(conditions, condition)
		conditionMap[fmt.Sprintf("%s-%d", gatewayName, port)] = conditions
	}
	expectedRouteStatus := akogatewayapitests.GetRouteStatusV1([]string{gatewayName}, namespace, ports, conditionMap)

	httpRoute, err := akogatewayapitests.GatewayClient.GatewayV1().HTTPRoutes(namespace).Get(context.TODO(), httpRouteName, metav1.GetOptions{})
	if err != nil || httpRoute == nil {
		t.Fatalf("Couldn't get the HTTPRoute, err: %+v", err)
	}
	akogatewayapitests.ValidateHTTPRouteStatus(t, &httpRoute.Status, &gatewayv1.HTTPRouteStatus{RouteStatus: *expectedRouteStatus})

	// Update
	rules[0].BackendRefs[0].Filters[0].RequestHeaderModifier.Set = nil
	akogatewayapitests.UpdateHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)

	g.Eventually(func() bool {
		httpRoute, err := akogatewayapitests.GatewayClient.GatewayV1().HTTPRoutes(namespace).Get(context.TODO(), httpRouteName, metav1.GetOptions{})
		if err != nil || httpRoute == nil {
			t.Logf("Couldn't get the HTTPRoute, err: %+v", err)
			return false
		}
		if len(httpRoute.Status.Parents) != len(ports) {
			return false
		}
		return apimeta.IsStatusConditionTrue(httpRoute.Status.Parents[0].Conditions, string(gatewayv1.GatewayConditionAccepted))
	}, 30*time.Second).Should(gomega.Equal(true))

	expectedRouteStatus = akogatewayapitests.GetRouteStatusV1([]string{gatewayName}, namespace, ports, conditionMap)

	httpRoute, err = akogatewayapitests.GatewayClient.GatewayV1().HTTPRoutes(namespace).Get(context.TODO(), httpRouteName, metav1.GetOptions{})
	if err != nil || httpRoute == nil {
		t.Fatalf("Couldn't get the HTTPRoute, err: %+v", err)
	}
	akogatewayapitests.ValidateHTTPRouteStatus(t, &httpRoute.Status, &gatewayv1.HTTPRouteStatus{RouteStatus: *expectedRouteStatus})
	akogatewayapitests.TeardownHTTPRoute(t, httpRouteName, namespace)
	akogatewayapitests.TeardownGateway(t, gatewayName, namespace)
	akogatewayapitests.TeardownGatewayClass(t, gatewayClassName)
}

func TestHTTPRouteGatewayWithEmptyHostnameInGateway(t *testing.T) {
	gatewayClassName := "gateway-class-04"
	gatewayName := "gateway-04"
	httpRouteName := "httproute-04"
	gwKey := "Gateway/" + DEFAULT_NAMESPACE + "/" + gatewayName
	gwClassKey := "GatewayClass/" + gatewayClassName
	namespace := "default"
	ports := []int32{8080}
	key := "HTTPRoute" + "/" + namespace + "/" + httpRouteName

	// gatewayclass
	akogatewayapiobjects.GatewayApiLister().UpdateGatewayClass(gatewayClassName, true)
	akogatewayapitests.SetupGatewayClass(t, gatewayClassName, akogatewayapilib.GatewayController)
	t.Logf("Created GatewayClass %s", gatewayClassName)
	waitAndverify(t, gwClassKey)

	// Gateway with empty hostname
	listeners := akogatewayapitests.GetListenersV1(ports, true)
	akogatewayapitests.SetupGateway(t, gatewayName, namespace, gatewayClassName, nil, listeners)
	t.Logf("Created Gateway %s without hostname", gatewayName)
	waitAndverify(t, gwKey)

	t.Logf("Now creating httproute")
	parentRefs := akogatewayapitests.GetParentReferencesV1([]string{gatewayName}, namespace, ports)
	hostnames := []gatewayv1.Hostname{"foo-8080.com"}
	akogatewayapitests.SetupHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)
	waitAndverify(t, key)

	// delete
	akogatewayapitests.TeardownHTTPRoute(t, httpRouteName, namespace)
	waitAndverify(t, key)
	akogatewayapitests.TeardownGateway(t, gatewayName, DEFAULT_NAMESPACE)
	waitAndverify(t, gwKey)
	akogatewayapitests.TeardownGatewayClass(t, gatewayClassName)
	waitAndverify(t, gwClassKey)
}

func TestHTTPRouteGatewayWithEmptyHostnameInHTTPRoute(t *testing.T) {
	gatewayClassName := "gateway-class-05"
	gatewayName := "gateway-05"
	httpRouteName := "httproute-05"
	gwKey := "Gateway/" + DEFAULT_NAMESPACE + "/" + gatewayName
	gwClassKey := "GatewayClass/" + gatewayClassName
	namespace := "default"
	ports := []int32{8080}
	key := "HTTPRoute" + "/" + namespace + "/" + httpRouteName

	// gatewayclass
	akogatewayapiobjects.GatewayApiLister().UpdateGatewayClass(gatewayClassName, true)
	akogatewayapitests.SetupGatewayClass(t, gatewayClassName, akogatewayapilib.GatewayController)
	t.Logf("Created GatewayClass %s", gatewayClassName)
	waitAndverify(t, gwClassKey)

	// Gateway
	listeners := akogatewayapitests.GetListenersV1(ports, false)
	akogatewayapitests.SetupGateway(t, gatewayName, namespace, gatewayClassName, nil, listeners)
	t.Logf("Created Gateway %s", gatewayName)
	waitAndverify(t, gwKey)

	// httproute without hostname
	t.Logf("Now creating httproute without hostname")
	parentRefs := akogatewayapitests.GetParentReferencesV1([]string{gatewayName}, namespace, ports)
	hostnames := []gatewayv1.Hostname{}
	akogatewayapitests.SetupHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)
	waitAndverify(t, key)

	// delete
	akogatewayapitests.TeardownHTTPRoute(t, httpRouteName, namespace)
	waitAndverify(t, key)
	akogatewayapitests.TeardownGateway(t, gatewayName, DEFAULT_NAMESPACE)
	waitAndverify(t, gwKey)
	akogatewayapitests.TeardownGatewayClass(t, gatewayClassName)
	waitAndverify(t, gwClassKey)
}

func TestHTTPRouteGatewayWithEmptyHostname(t *testing.T) {
	gatewayClassName := "gateway-class-06"
	gatewayName := "gateway-06"
	httpRouteName := "httproute-06"
	gwKey := "Gateway/" + DEFAULT_NAMESPACE + "/" + gatewayName
	gwClassKey := "GatewayClass/" + gatewayClassName
	namespace := "default"
	ports := []int32{8080}
	key := "HTTPRoute" + "/" + namespace + "/" + httpRouteName

	// gatewayclass
	akogatewayapiobjects.GatewayApiLister().UpdateGatewayClass(gatewayClassName, true)
	akogatewayapitests.SetupGatewayClass(t, gatewayClassName, akogatewayapilib.GatewayController)
	t.Logf("Created GatewayClass %s", gatewayClassName)
	waitAndverify(t, gwClassKey)

	// Gateway without hostname
	listeners := akogatewayapitests.GetListenersV1(ports, true)
	akogatewayapitests.SetupGateway(t, gatewayName, namespace, gatewayClassName, nil, listeners)
	t.Logf("Created Gateway %s", gatewayName)
	waitAndverify(t, gwKey)

	// httproute without hostname
	t.Logf("Now creating httproute without hostname")
	parentRefs := akogatewayapitests.GetParentReferencesV1([]string{gatewayName}, namespace, ports)
	hostnames := []gatewayv1.Hostname{}
	akogatewayapitests.SetupHTTPRoute(t, httpRouteName, namespace, parentRefs, hostnames, nil)
	waitAndverify(t, key)

	// delete
	akogatewayapitests.TeardownHTTPRoute(t, httpRouteName, namespace)
	waitAndverify(t, key)
	akogatewayapitests.TeardownGateway(t, gatewayName, DEFAULT_NAMESPACE)
	waitAndverify(t, gwKey)
	akogatewayapitests.TeardownGatewayClass(t, gatewayClassName)
	waitAndverify(t, gwClassKey)
}

