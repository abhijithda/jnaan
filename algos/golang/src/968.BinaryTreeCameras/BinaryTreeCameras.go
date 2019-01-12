package binarytreecameras

import (
	"fmt"
	"internal/tree"
	"log"
	"os"
)

func init() {
	const logFile = "log.txt"
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	// defer f.Close()
	log.SetOutput(f)
}

// TreeNode is node of a binary tree node.
type TreeNode = tree.Node

func install(node, p, pp *TreeNode, covered map[*TreeNode]bool) int {
	log.Printf("Given node: %v; p: %v; pp: %v; covered: %v",
		node, p, pp, covered)
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		cnt := 0
		if !covered[node] {
			// Install camera to node's parent, which covers itself and it's
			// sibling, along with node's parent's parent.
			cnt++
			if p != nil {
				covered[p] = true
				if p.Left != nil {
					covered[p.Left] = true
				}
				if p.Right != nil {
					covered[p.Right] = true
				}
				if pp != nil {
					covered[pp] = true
				}
			}
		}
		log.Printf("Returning cnt: %d for leaf node: %v!", cnt, node)
		return cnt
	}
	cnt := 0
	if node.Left != nil {
		cnt += install(node.Left, node, p, covered)
	}

	if node.Right != nil {
		cnt += install(node.Right, node, p, covered)
	}

	if !covered[node] {
		// Install camera to node's parent.
		cnt++

		// Installing camera to node's parent covers node itself and it's
		// sibling, along with node's parent's parent.
		if p != nil {
			covered[p] = true
			if pp != nil {
				covered[pp] = true
			}
			if p.Left != nil {
				covered[p.Left] = true
			}
			if p.Right != nil {
				covered[p.Right] = true
			}
		}
	}
	log.Printf("Returning cnt: %d for leaf node: %v with coverage: %v!",
		cnt, node, covered)
	return cnt
}

func minCameraCover(root *TreeNode) int {
	return install(root, nil, nil, map[*TreeNode]bool{})
}

func installCamera(node *TreeNode, level int) int {
	log.Printf("Given node: %v; level: %d", node, level)
	if node == nil {
		return 0
	}
	cnt := 0

	if level%3 == 1 {
		cnt++
	} else if level%3 == 0 {
		if node.Left == nil && node.Right == nil {
			// If there a no child node for installing camera, then install
			//	camera here itself as previous node also didn't had camera.
			cnt++
			log.Printf("Incremented count (%d) for leaf node: %v", cnt, node)
		} else {
			// If there are children, then see whether installing to current
			//  node, or installing to children would result in lesser cameras.
			nirc := installCamera(node.Left, 1) + installCamera(node.Right, 1)
			log.Println("No camera on cur node count:", nirc)
			irc := installCamera(node, 1)
			log.Println("Camera on cur node count:", irc)

			if nirc < irc {
				cnt += nirc
			} else {
				cnt += irc
			}
		}
		return cnt
	}
	if node.Left != nil {
		cnt += installCamera(node.Left, level+1)
	}
	if node.Right != nil {
		cnt += installCamera(node.Right, level+1)
	}
	log.Printf("For node: %v; level: %d; Count: %d", node, level, cnt)
	return cnt
}

func minCameraCoverTake2(root *TreeNode) int {
	log.Println("\n\n\nGiven Root:", root)
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	nirc := installCamera(root, 0)
	log.Println("No camera on root count:", nirc)
	irc := installCamera(root, 1)
	log.Println("Camera on root count:", irc)

	if nirc < irc {
		return nirc
	}
	return irc
}

func installCameraOld(node *TreeNode, covered int, preinstall bool) int {
	log.Printf("Node: %v, covered: %d; preinstall: %v", node, covered, preinstall)
	if node == nil {
		return 0
	}
	count := 0
	// covered++
	if covered == 1 {
		// If preinstalled (i.e., installed in 3rd node, then don't install
		//  new one, instead just move 3rd camera to 4th (1) node.),
		// 	then no need to increment counter.
		if !preinstall {
			count++
		}
		preinstall = false
	} else if covered == 3 {
		preinstall = true
		count++
		covered = 0
	}
	if node.Left != nil {
		count += installCameraOld(node.Left, covered+1, preinstall)
	}
	if node.Right != nil {
		count += installCameraOld(node.Right, covered+1, preinstall)
	}
	log.Println("Count:", count)
	return count
}

func minCameraCoverTake1(root *TreeNode) int {
	log.Println("Given root:", root)
	if root == nil {
		return 0
	}
	rootc := installCameraOld(root, 1, false)
	rootnc := installCameraOld(root, 0, false)

	count := rootc
	if rootnc < rootc && rootnc != 0 {
		count = rootnc
	}
	return count
}
