package schema

var logger string = `
[id]
	type = id
	fid  = 1
[pinyin]
	nid  = 2
[partial]
	fid  = 3
[total]
	type=numeric
	index=self
	fid  = 4
[lastnum]
	type=numeric
	index=self
	fid  = 5
[currnum]
	type=numeric
	index=self
	fid  = 6
[currtag]
	fid = 7
[body]
	type=body
	fid = 8
`

var logger1 string = `
[fields]

[fields.id]
	type = "id"
	fid  = 1
[fields.pinyin]
	nid  = 2
[fields.partial]
	fid  = 3
[fields.total]
	type="numeric"
	index="self"
	fid  = 4
[fields.lastnum]
	type="numeric"
	index="self"
	fid  = 5
[fields.currnum]
	type="numeric"
	index="self"
	fid  = 6
[fields.currtag]
	fid = 7
[fields.body]
	type="body"
	fid = 8
`
