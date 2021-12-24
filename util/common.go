package util

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"ftk8s/base/enum"
	"ftk8s/model"

	"github.com/google/uuid"
)

func IsSupportResourceKind(resourceKind string) bool {
	flag := false
	for _, v := range enum.SupportResourceKindSlice {
		if resourceKind == v {
			flag = true
			return flag
		}
	}
	return flag
}

// GenerateUUIDV4
func GenerateUUIDV4(prefix string) (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed to generate uuid v4, error message: %w", err)
	}
	uuidStringTemp := uuidObj.String()
	uuidString := strings.Replace(uuidStringTemp, "-", "", -1)
	result := fmt.Sprint(prefix, uuidString)

	return result, err
}

// GetLenOfSyncMap get length of sync.Map
func GetLenOfSyncMap(m *sync.Map) int {
	length := 0
	m.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	return length
}

// ByteSliceToString converts slice to string without copy.
// Use at your own risk.
func ByteSliceToString(b []byte) (s string) {
	if len(b) == 0 {
		return ""
	}
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

// StringToByteSlice converts string to slice without copy.
// Use at your own risk.
func StringToByteSlice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

func GenerateWebpermSlice(id string, oriNodes []*model.Webperm, resNodes *[]*model.Webperm) {
	children := FindChildWebperm(id, oriNodes)
	for _, child := range children {
		*resNodes = append(*resNodes, child)
		if IsHasChildWebperm(child.ID, oriNodes) {
			GenerateWebpermSlice(child.ID, oriNodes, resNodes)
		}
	}
}

func FindChildWebperm(id string, nodes []*model.Webperm) []*model.Webperm {
	children := make([]*model.Webperm, 0)
	for _, v := range nodes {
		if id == v.ParentID {
			children = append(children, v)
		}
	}
	return children
}

func IsHasChildWebperm(id string, nodes []*model.Webperm) bool {
	flag := false
	for _, v := range nodes {
		if id == v.ParentID {
			flag = true
			break
		}
	}
	return flag
}

func GenerateTreeWebpermTree(node *model.RespWebpermTree, nodes []*model.RespWebpermTree) {
	children := FindChildWebpermTree(node, nodes)
	for _, child := range children {
		node.Children = append(node.Children, child)
		if IsHasChildWebpermTree(child, nodes) {
			GenerateTreeWebpermTree(child, nodes)
		}
	}
}

func FindChildWebpermTree(node *model.RespWebpermTree, nodes []*model.RespWebpermTree) []*model.RespWebpermTree {
	children := make([]*model.RespWebpermTree, 0)
	for _, v := range nodes {
		if node.ID == v.ParentID {
			children = append(children, v)
		}
	}
	return children
}

func IsHasChildWebpermTree(node *model.RespWebpermTree, nodes []*model.RespWebpermTree) bool {
	flag := false
	for _, v := range nodes {
		if node.ID == v.ParentID {
			flag = true
			break
		}
	}
	return flag
}

func GenerateTreeWebpermTreeByRoleOnlyDisplay(node *model.RespWebpermTreeByRoleOnlyDisplay, nodes []*model.RespWebpermTreeByRoleOnlyDisplay) {
	children := FindChildWebpermTreeByRoleOnlyDisplay(node, nodes)
	for _, child := range children {
		node.Children = append(node.Children, child)
		if IsHasChildWebpermTreeByRoleOnlyDisplay(child, nodes) {
			GenerateTreeWebpermTreeByRoleOnlyDisplay(child, nodes)
		}
	}
}

func FindChildWebpermTreeByRoleOnlyDisplay(node *model.RespWebpermTreeByRoleOnlyDisplay, nodes []*model.RespWebpermTreeByRoleOnlyDisplay) []*model.RespWebpermTreeByRoleOnlyDisplay {
	children := make([]*model.RespWebpermTreeByRoleOnlyDisplay, 0)
	for _, v := range nodes {
		if node.ID == v.ParentID {
			children = append(children, v)
		}
	}
	return children
}

func IsHasChildWebpermTreeByRoleOnlyDisplay(node *model.RespWebpermTreeByRoleOnlyDisplay, nodes []*model.RespWebpermTreeByRoleOnlyDisplay) bool {
	flag := false
	for _, v := range nodes {
		if node.ID == v.ParentID {
			flag = true
			break
		}
	}
	return flag
}

func GenerateTreeRoleWebpermRightsTree(node *model.RespWebpermRights, nodes []*model.RespWebpermRights) {
	children := FindChildRoleWebpermRightsTree(node, nodes)
	for _, child := range children {
		node.Children = append(node.Children, child)
		if IsHasChildRoleWebpermRightsTree(child, nodes) {
			GenerateTreeRoleWebpermRightsTree(child, nodes)
		}
	}
}

func FindChildRoleWebpermRightsTree(node *model.RespWebpermRights, nodes []*model.RespWebpermRights) []*model.RespWebpermRights {
	children := make([]*model.RespWebpermRights, 0)
	for _, v := range nodes {
		if node.ID == v.ParentID {
			children = append(children, v)
		}
	}
	return children
}

func IsHasChildRoleWebpermRightsTree(node *model.RespWebpermRights, nodes []*model.RespWebpermRights) bool {
	flag := false
	for _, v := range nodes {
		if node.ID == v.ParentID {
			flag = true
			break
		}
	}
	return flag
}

// PreTraversalRoleWebpermRights preorder traversal of the tree
func PreTraversalRoleWebpermRights(root *model.RespWebpermRights) []*model.RespWebpermRights {
	if root == nil {
		return nil
	}
	nodeSlice := make([]*model.RespWebpermRights, 0)
	stack := []*model.RespWebpermRights{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for i := len(node.Children) - 1; i > -1; i-- {
			stack = append(stack, node.Children[i])
		}
		data := node
		data.Children = make([]*model.RespWebpermRights, 0)
		nodeSlice = append(nodeSlice, data)
	}

	return nodeSlice
}
