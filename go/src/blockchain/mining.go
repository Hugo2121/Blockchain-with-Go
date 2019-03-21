package blockchain

import (
	"work_queue"
	// "fmt"
)

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	start uint64
	end uint64
	blk Block
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

func (worker miningWorker) Run() interface{} {
	for i := worker.start; i <= worker.end; i++ {
		worker.blk.Proof = i
		// worker.blk.Hash = worker.blk.CalcHash()
		// fmt.Println(i)
		if worker.blk.ValidHash() {
			return MiningResult{i, true}
		}
	}
	return MiningResult{0, false}
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	// TODO
	eachLen := (end - start) / chunks
	remainder := (end - start) % chunks
	q := work_queue.Create(uint(workers), uint(chunks))
	// fmt.Println("eachLen is ",eachLen)
	// fmt.Println("start is ", start)
	// fmt.Println("end is ", end)
	// fmt.Println(chunks)
	count := 0

	for i := start; i <= end; i += eachLen {
		count++
		q.Enqueue(miningWorker{i, i+eachLen, blk})
		// fmt.Println(i, i+eachLen)
		if uint64(count) == chunks {
			break
		}
	}
	ret := MiningResult{}
	for j := uint64(0); j < chunks; j++ {
		ret = (<-q.Results).(MiningResult)
		if ret.Found {
			q.Shutdown()
			break
		}
	}

	if !ret.Found {
 		q.Enqueue(miningWorker{end-remainder, end, blk})
 		ret = (<-q.Results).(MiningResult)
 		q.Shutdown()
	}



	return ret

}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}

