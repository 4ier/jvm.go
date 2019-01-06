package main

import (
	"fmt"
	"strings"

	"jvm.go/classfile"
	"jvm.go/classpath"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)

	if err != nil {
		panic(err)
	}

	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlag())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceName())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("    %s\n", f.Name())
	}

	fmt.Printf("methods count: %v\n", len(cf.Methods()))

	for _, m := range cf.Methods() {
		fmt.Printf("    %s\n", m.Name())
	}

}
