Initial Golang implementation of the "newscast model is a general approach for communication
in large agent-based distributed systems"

Named for Canada's premiere newscaster, [Peter Mansbridge](https://en.wikipedia.org/wiki/Peter_Mansbridge) 

This is a simplied implementation of the protocol outlined in the [original paper](http://www.cs.unibo.it/bison/publications/ap2pc03.pdf), based
on the author's subsequent paper on their own [implementation](http://www.soc.napier.ac.uk/~benp/dream/dreampaper17.pdf)

Will hopefully expand to cover more of the details outlined in the original paper.

### Usage:

* implement the ```Agent``` interface
* create a Correspondent by calling ```NewCorrespondent``` and passing it a WireService implementation, Agent implementation, and other settings
  
  ```go 
     encoder := wire.GobWireEncoder{}
	 wireService := wire.NewUdpWireService(*port, encoder)
     c := correspondent.NewCorrespondent(TestAgent{id: agentId}, wireService, *delay, *seed, *cacheSize)
  ```

* call ```Correspondent.StartReporting()``` this is a blocking call, so it is recommended to do this in a goroutine
