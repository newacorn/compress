package zstd

func (d *Decoder) NilSrc() {
	d.syncStream.br.r = nil
}
