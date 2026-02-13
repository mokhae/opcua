package attrs

import "github.com/mokhae/opcua/ua"

func BrowseName(name string) *ua.QualifiedName {
	return &ua.QualifiedName{Name: name}
}

func DisplayName(name, locale string) *ua.LocalizedText {
	l := &ua.LocalizedText{Text: name, Locale: locale}
	l.UpdateMask()
	return l
}

func InverseName(name, locale string) *ua.LocalizedText {
	l := &ua.LocalizedText{Text: name, Locale: locale}
	l.UpdateMask()
	return l
}

func Description(text, locale string) *ua.LocalizedText {
	l := &ua.LocalizedText{Text: text, Locale: locale}
	l.UpdateMask()
	return l
}

func NodeClass(n ua.NodeClass) int32 {
	return int32(n)
}

func DataType(id *ua.NodeID) *ua.ExpandedNodeID {
	return &ua.ExpandedNodeID{NodeID: id}
}
