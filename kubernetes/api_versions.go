package kubernetes

import (
	"encoding/json"
	"log"
	"strconv"

	kubernetes "k8s.io/client-go/kubernetes"
)

// Convert between two types by converting to/from JSON. Intended to switch
// between multiple API versions, as they are strict supersets of one another.
// item and out are pointers to structs
func Convert(item, out interface{}) error {
	bytes, err := json.Marshal(item)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, out)
	if err != nil {
		return err
	}

	// Converting between maps and structs only occurs when autogenerated resources convert the result
	// of an HTTP request. Those results do not contain omitted fields, so no need to set them.
	// if _, ok := item.(map[string]interface{}); !ok {
	// 	setOmittedFields(item, out)
	// }

	return nil
}

// ServerVersionPre1_9 reads the Kubernetes API verions and returns true if less
// than v1.9
func ServerVersionPre1_9(conn *kubernetes.Clientset) bool {
	ver, _ := conn.ServerVersion()
	minor, _ := strconv.Atoi(string(ver.Minor[0]))
	log.Printf("[INFO] Kubernetes Server version: %#v", ver)

	if ver.Major == "1" && minor < 9 {
		return true
	}

	return false
}
