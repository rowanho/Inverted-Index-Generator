package invertedindex

// InvertedIndexEntry contains the term followed by the
// number of times it has appeared across all documents
// and an array of documents it is persent in
type InvertedIndexEntry struct {
	Term            uint64
	Frequency       int
	DocumentListing []int
}

// InvertedIndex contains a hash map to easily check if the
// term is present and an array of InvertedIndexEntry
type InvertedIndex struct {
	HashMap map[uint64]*InvertedIndexEntry
	Items   []*InvertedIndexEntry
}

// FindItem returns the position of a given
// Item in an Inverted Index
func (invertedIndex *InvertedIndex) FindItem(Term uint64) []int {
	indexes := make([]int, 0)
	for index, item := range invertedIndex.Items {
		if item.Term == Term {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

// AddItem works by first checking if a given term is already present
// in the inverse index or not by checking the hashmap. If it is
// present it updates the Items by increasing the frequency and
// adding the document it is found in. If it is not present it
// adds it to the hash map and adds it to the items list
func (invertedIndex *InvertedIndex) AddItem(Term uint64, Document int) {
	if invertedIndex.HashMap[Term] != nil {
		// log.Println("Index item", Term, "already exists :: updating existing item")

		FoundItemPosition := invertedIndex.FindItem(Term)

		invertedIndex.Items[FoundItemPosition].Frequency++
		invertedIndex.Items[FoundItemPosition].DocumentListing = append(invertedIndex.Items[FoundItemPosition].DocumentListing, Document)
	} else {
		// log.Println("Index item", Term, " does not exist :: creating new item")

		InvertedIndexEntry := &InvertedIndexEntry{
			Term:            Term,
			Frequency:       1,
			DocumentListing: []int{Document},
		}

		invertedIndex.HashMap[Term] = InvertedIndexEntry
		invertedIndex.Items = append(invertedIndex.Items, InvertedIndexEntry)
	}
}

// CreateInvertedIndex initializes an
// empty Inverted Index
func CreateInvertedIndex() *InvertedIndex {
	invertedIndex := &InvertedIndex{
		HashMap: make(map[uint64]*InvertedIndexEntry),
		Items:   []*InvertedIndexEntry{},
	}
	return invertedIndex
}


func GenerateFpMap(fps []uint64) map[uint64]bool {
	fpMap := make(map[uint64]bool)

	for _, fp := range fps {
		if _, value := fpMap[fp]; !value {
			fpMap[fp] = true
		}
	}

	return fpMap
}

// GenerateInvertedIndex for each document list
// gets each word as a token, processes it and
// generates a hash map for each document
// using them it then generates the
// inverted index of all words
func GenerateInvertedIndex(fpList [][]uint64) InvertedIndex {
	fpMaps := make([]map[uint64]bool, len(fpList))

	for _, fps := range fpList {
		fpMap := GenerateFpMap(fps)
		fpMaps = append(fpMaps, fpMap)
	}

	// Create an empty inverted index
	invertedIndex := CreateInvertedIndex()

	// Using the generated hash maps add
	// each word to the inverted index
	for index, fpMap := range fpMaps {
		for fp := range fpMap {
			invertedIndex.AddItem(fp, index)
		}
	}
	return *invertedIndex
}