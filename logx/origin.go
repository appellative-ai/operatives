package logx

import (
	"errors"
	"fmt"
	"strings"
)

const (
	originNameFmt  = "%v:%v:%v"
	serviceKind    = "service"
	RegionKey      = "region"
	ZoneKey        = "zone"
	SubZoneKey     = "sub-zone"
	HostKey        = "host"
	InstanceIdKey  = "instance-id"
	ServiceNameKey = "service-name"
	CollectiveKey  = "collective"
	DomainKey      = "domain"
)

// OriginT - location
type OriginT struct {
	Name        string `json:"name"`
	Region      string `json:"region"`
	Zone        string `json:"zone"`
	SubZone     string `json:"sub-zone"`
	Host        string `json:"host"`
	ServiceName string `json:"service-name"`
	InstanceId  string `json:"instance-id"`
	Collective  string `json:"collective"`
	Domain      string `json:"domain"`
}

func (o *OriginT) String() string { return o.Name }

func (o *OriginT) IsLocalCollective(name string) bool {
	if strings.HasPrefix(name, o.Collective+":") {
		return true
	}
	return false
}

/*
func NewOriginFromMessage(m *Message, collective, domain string) (OriginT, *Status) {
	cfg, status := MapContent(m)
	if !status.OK() {
		return OriginT{}, status
	}
	return NewOrigin(cfg, collective, domain)
}


*/

func newOrigin(origin *OriginT, m map[string]string) error {
	//var origin OriginT

	if m == nil {
		return errors.New("origin map is nil")
	}

	origin.Collective = m[CollectiveKey]
	if origin.Collective == "" {
		return errors.New(fmt.Sprintf("config map does not contain key: %v", CollectiveKey))
	}
	origin.Domain = m[DomainKey]
	if origin.Domain == "" {
		return errors.New(fmt.Sprintf("config map does not contain key: %v", DomainKey))
	}
	origin.Region = m[RegionKey]
	if origin.Region == "" {
		return errors.New(fmt.Sprintf("config map does not contain key: %v", RegionKey))
	}
	origin.Zone = m[ZoneKey]
	if origin.Zone == "" {
		return errors.New(fmt.Sprintf("config map does not contain key: %v", ZoneKey))
	}
	origin.Host = m[HostKey]
	if origin.Host == "" {
		return errors.New(fmt.Sprintf("config map does not contain key: %v", HostKey))
	}

	origin.ServiceName = m[ServiceNameKey]
	if origin.ServiceName == "" {
		origin.ServiceName = origin.Host
	}
	origin.SubZone = m[SubZoneKey]
	origin.InstanceId = m[InstanceIdKey]
	origin.Name = name(origin)
	return nil
}

func name(o *OriginT) string {
	var name1 = fmt.Sprintf(originNameFmt, o.Collective, o.Domain, serviceKind)

	if o.Region != "" {
		name1 += "/" + o.Region
	}
	if o.Zone != "" {
		name1 += "/" + o.Zone
	}
	if o.SubZone != "" {
		name1 += "/" + o.SubZone
	}
	if o.ServiceName != "" {
		name1 += "/" + o.ServiceName
	}
	if o.InstanceId != "" {
		name1 += "#" + o.InstanceId
	}
	return name1
}

/*
//if o.SubZone == "" {
	//	messaging.Reply(m, messaging.ConfigMapContentError(nil, SubZoneKey), NamespaceName)
	//	return
	//}

	if o.InstanceId == "" {
		messaging.Reply(m, messaging.ConfigMapContentError(a, InstanceIdKey), a.Name())
		return
	}



*/

/*
func (o Origin) Uri(class string) string {
	return fmt.Sprintf(uriFmt, class, o)
}

func (o Origin) String() string { return "" }


*/

/*

var uri = o.Region

if o.Zone != "" {
uri += "." + o.Zone
}
if o.SubZone != "" {
uri += "." + o.SubZone
}
if o.Host != "" {
uri += "." + o.Host
}
return uri
}


*/
