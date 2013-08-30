// Copyright (c) 2013 - Michael Woolnough <michael.woolnough@gmail.com>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package Minecraft will be a full featured minecraft level editor/viewer.
package minecraft

var (
	transparentBlocks []uint16
	lightBlocks       map[uint8]uint8
)

func init() {
	transparentBlocks = []uint16{0, 6, 18, 20, 26, 27, 28, 29, 30, 31, 33, 34, 37, 38, 39, 40, 50, 51, 52, 54, 55, 59, 63, 64, 65, 66, 69, 70, 71, 75, 76, 77, 78, 79, 81, 83, 85, 90, 92, 93, 94, 96, 102, 106, 107, 117, 118, 119, 120, 750}
	lightBlocks = map[uint8]uint8{
		10:  15,
		11:  15,
		39:  1,
		50:  14,
		51:  15,
		62:  13,
		74:  13,
		76:  7,
		89:  15,
		90:  11,
		91:  15,
		94:  9,
		95:  15,
		117: 1,
		119: 15,
		120: 1,
		122: 1,
		124: 16,
		130: 15,
		138: 15,
	}
}
