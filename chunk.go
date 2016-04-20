package main

type ChunkReq struct {
	msglen   uint32
	msgcmd   uint32
	msgseq   uint32
	nodeid   [6]byte
	fid      [16]byte
	datasrc  uint32
	blockcur uint32
	mapbegin uint32
	mapsize  uint32
	bitmap   []byte
}

type ChunkResp struct {
	msglen  uint32
	msgcmd  uint32
	msgseq  uint32
	nodeid  [6]byte
	fid     [16]byte
	status  uint32
	datasrc uint32
	index   uint32
	length  uint32
	data    []byte
}

type Chunk_One_Req struct {
	msglen    uint32
	msgcmd    uint32
	msgseq    uint32
	filesize  uint64
	fid       [16]byte
	datasrc   uint32
	blockcur  uint32
	mapbegin  uint32
	mapsize   uint32
	downbegin uint32
	bitmap    []byte
}

type Chunk_NoExist struct {
	fid   []byte
	index uint32
}
