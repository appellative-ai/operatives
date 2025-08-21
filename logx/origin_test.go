package logx

import "fmt"

func ExampleNewOrigin() {
	m := map[string]string{
		RegionKey:      "region",
		ZoneKey:        "zone",
		SubZoneKey:     "sub-zone",
		HostKey:        "host",
		ServiceNameKey: "service-name",
		InstanceIdKey:  "instance-id",
		CollectiveKey:  "collective",
		DomainKey:      "domain",
	}
	var o OriginT
	status := newOrigin(&o, m)
	fmt.Printf("test: newOrigin() -> [%v] [status:%v]\n", o, status)

	o.Zone = ""
	fmt.Printf("test: Name() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = ""
	fmt.Printf("test: Name() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = "sub-zone"
	o.Host = ""
	fmt.Printf("test: Name() -> [%v]\n", o)

	o.Zone = "zone"
	o.SubZone = "sub-zone"
	o.Host = "host"
	o.InstanceId = "instance-id"
	fmt.Printf("test: Name() -> [%v]\n", o)

	//Output:
	//test: newOrigin() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region zone sub-zone host service-name instance-id collective domain}] [status:<nil>]
	//test: Name() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region  sub-zone host service-name instance-id collective domain}]
	//test: Name() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region zone  host service-name instance-id collective domain}]
	//test: Name() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region zone sub-zone  service-name instance-id collective domain}]
	//test: Name() -> [{collective:domain:service/region/zone/sub-zone/service-name#instance-id region zone sub-zone host service-name instance-id collective domain}]

}

func ExampleNewOrigin_Error() {
	m := make(map[string]string)

	var o OriginT
	status := newOrigin(&o, m)
	fmt.Printf("test: newOrigin() -> [%v] [status:%v]\n", o, status)

	m[CollectiveKey] = "collective"
	//var o1 OriginT
	status = newOrigin(&o, m)
	fmt.Printf("test: newOrigin() -> [%v] [status:%v]\n", o, status)

	m[DomainKey] = "domain"
	//var o2 OriginT
	status = newOrigin(&o, m)
	fmt.Printf("test: newOrigin() -> [%v] [status:%v]\n", o, status)

	m[RegionKey] = "region"
	//var o3 OriginT
	status = newOrigin(&o, m)
	fmt.Printf("test: newOrigin() -> [%v] [status:%v]\n", o, status)

	m[ZoneKey] = "zone"
	//var o4 OriginT
	status = newOrigin(&o, m)
	fmt.Printf("test: newOrigin() -> [%v] [status:%v]\n", o, status)

	m[HostKey] = "host"
	//var o5 OriginT
	status = newOrigin(&o, m)
	fmt.Printf("test: newOrigin() -> [%v] [status:%v]\n", o, status)

	//Output:
	//test: newOrigin() -> [{        }] [status:config map does not contain key: collective]
	//test: newOrigin() -> [{       collective }] [status:config map does not contain key: domain]
	//test: newOrigin() -> [{       collective domain}] [status:config map does not contain key: region]
	//test: newOrigin() -> [{ region      collective domain}] [status:config map does not contain key: zone]
	//test: newOrigin() -> [{ region zone     collective domain}] [status:config map does not contain key: host]
	//test: newOrigin() -> [{collective:domain:service/region/zone/host region zone  host host  collective domain}] [status:<nil>]

}

func ExampleIsLocalCollectiveNewOrigin() {
	m := map[string]string{
		RegionKey:      "region",
		ZoneKey:        "zone",
		SubZoneKey:     "sub-zone",
		HostKey:        "host",
		ServiceNameKey: "service-name",
		InstanceIdKey:  "instance-id",
		CollectiveKey:  "collective",
		DomainKey:      "domain",
	}
	var o OriginT
	newOrigin(&o, m)

	name1 := ""
	fmt.Printf("test: IsLocalCollective(\"%v\") -> [local:%v]\n", name1, o.IsLocalCollective(name1))

	name1 = o.Collective
	fmt.Printf("test: IsLocalCollective(\"%v\") -> [local:%v]\n", name1, o.IsLocalCollective(name1))

	name1 = o.Collective + ":"
	fmt.Printf("test: IsLocalCollective(\"%v\") -> [local:%v]\n", name1, o.IsLocalCollective(name1))

	//Output:
	//test: IsLocalCollective("") -> [local:false]
	//test: IsLocalCollective("collective") -> [local:false]
	//test: IsLocalCollective("collective:") -> [local:true]

}
