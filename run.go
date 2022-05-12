package main

import (
    "fmt"
    "github.com/terra-farm/go-xen-api-client"
)

const XEN_API_URL string = "https://192.168.122.225"
const XEN_API_USERNAME string = "root"
const XEN_API_PASSWORD string = "aspirine"
const VM_NAME_LABEL = "test"

func main() {
    xapi, err := xenapi.NewClient(XEN_API_URL, nil)
    if err != nil {
        panic(err)
    }

    session, err := xapi.Session.LoginWithPassword(XEN_API_USERNAME, XEN_API_PASSWORD, "1.0", "example")
    if err != nil {
        panic(err)
    }

    vms, err := xapi.VM.GetByNameLabel(session, VM_NAME_LABEL)
    if err != nil {
        panic(err)
    }

    if len(vms) == 0 {
        panic(fmt.Errorf("No VM template with name label %q has been found", VM_NAME_LABEL))
    }

    if len(vms) > 1 {
        panic(fmt.Errorf("More than one VM with name label %q has been found", VM_NAME_LABEL))
    }

    vm := vms[0]

    xapi.VM.Start(session, vm, false, false)
    if err != nil {
        panic(err)
    }

    err = xapi.Session.Logout(session)
    if err != nil {
        panic(err)
    }
}