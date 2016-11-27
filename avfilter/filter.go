// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import "gopkg.in/targodan/ffgopeg.v1/avutil"

// GetFilterByName gets a filter definition matching the given name.
//
// C-Function: avfilter_get_by_name
func GetFilterByName(n string) *Filter {
	return (*Filter)(C.avfilter_get_by_name(C.CString(n)))
}

// Register registers a filter.
//
// C-Function: avfilter_register
func (f *Filter) Register() avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avfilter_register((*C.struct_AVFilter)(f))))
}

func filterNext(f *Filter) *Filter {
	return (*Filter)(C.avfilter_next((*C.struct_AVFilter)(f)))
}

// Filters returns a channel which can be used to iterate over the filters.
//
// Usage:
//
//     for f := range avdevice.Filters() {
//         // ...
//     }
//
// C-Function: avfilter_next
func Filters() <-chan *Filter {
	ch := make(chan *Filter)

	var e *Filter
	go func() {
		for {
			e = filterNext(e)
			if e == nil {
				break
			}
			ch <- e
		}
		close(ch)
	}()

	return ch
}
