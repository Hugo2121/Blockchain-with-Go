package main

import (
	//"exer10"
	"fmt"
	// "math"
	 "work_queue"
	 "blockchain"
	 "encoding/hex"
)

func main() {
	//fi := exer10.Fibonacci(30)
	//fmt.Println(fi)

	// pt := exer10.NewPoint(3, 4)
	// exer10.TurnDouble(&pt, 3*math.Pi/2)

	// fmt.Println(pt)

	// a := exer10.NewPoint(1,2)
	// b := exer10.NewPoint(-3,4)
	// c := exer10.NewPoint(5,-6)

	// tri := exer10.Triangle{a,b,c}
	// exer10.TurnDouble(&tri, math.Pi)
	// fmt.Println(tri)

	// exer10.DrawCircle(100, 100, "src/exer10/out.png")

	 q := work_queue.Create(2,4)
	 fmt.Println(q)

	 a := blockchain.Initial(16)
	 fmt.Println(a.ValidHash())
	 a.SetProof(56231)
	 fmt.Println(a.ValidHash())

	 //b0 := blockchain.Initial(19)
	 //fmt.Println(b0.ValidHash())
	 //b0.SetProof(87745)
	 //fmt.Println(b0.ValidHash())
	 //b1 := b0.Next("hash example 1234")
	 //b1.SetProof(1407891)
	 //fmt.Println(b1.ValidHash())
	 //b1.SetProof(346082)
	 //fmt.Println(b1.ValidHash())


	 b0 := blockchain.Initial(20)
	 b0.Mine(1)
	 fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	 b1 := b0.Next("this is an interesting message")
	 b1.Mine(1)
	 fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	 b2 := b1.Next("this is not interesting")
	 b2.Mine(1)
	 fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))

	 chain := blockchain.Blockchain{}
	 chain.Add(b0)
	 chain.Add(b1)
	 chain.Add(b2)
	 fmt.Println(chain.IsValid())
}
