package main

type NodeContent = interface{}

type Node struct {
	Data NodeContent
	next *Node
	prev *Node
}
