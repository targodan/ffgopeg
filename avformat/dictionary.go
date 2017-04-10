// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avformat

//#cgo pkg-config: libavutil
//#include <libavutil/dict.h>
import "C"

// Count returns the number of entries in dictionary.
//
// C-Function: av_dict_count
func (d *Dictionary) Count() int {
	return int(C.av_dict_count((*C.struct_AVDictionary)(d)))
}

// Get returns a dictionary entry with matching key.
//
// C-Function: av_dict_count
func (d *Dictionary) Get(key string) *DictionaryEntry {
	var entry *DictionaryEntry
	entry = (*DictionaryEntry)(C.av_dict_get((*C.struct_AVDictionary)(d), C.CString(key), (*C.struct_AVDictionaryEntry)(entry), C.AV_DICT_MATCH_CASE))

	return entry
}

// Next returns the entry in the dictionary that follows the provided entry.
func (d *Dictionary) Next(prev *DictionaryEntry) *DictionaryEntry {
	return (*DictionaryEntry)(C.av_dict_get((*C.struct_AVDictionary)(d), C.CString(""), (*C.struct_AVDictionaryEntry)(prev), C.AV_DICT_IGNORE_SUFFIX))
}

// Key returns the entry key.
//
// C-Variable: AVDictionaryEntry::key
func (de *DictionaryEntry) Key() string {
	return C.GoString(de.key)
}

// Value returns the entry value.
//
// C-Variable: AVDictionaryEntry::value
func (de *DictionaryEntry) Value() string {
	return C.GoString(de.value)
}
