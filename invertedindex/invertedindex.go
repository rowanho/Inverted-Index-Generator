package invertedindex

// InvertedIndexEntry contains the term followed by the
// number of times it has appeared across all documents
// and an array of documents it is persent in
type InvertedIndexEntry struct {
	Term            uint64
	Frequencies       map[int]int
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
func (invertedIndex *InvertedIndex) FindItem(Term uint64) int {
	for index, item := range invertedIndex.Items {
		if item.Term == Term {
			return index
		}
	}
	
	panic("Not Found")
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

		invertedIndex.Items[FoundItemPosition].Frequencies[Document]++
		invertedIndex.Items[FoundItemPosition].DocumentListing = append(invertedIndex.Items[FoundItemPosition].DocumentListing, Document)
	} else {
		// log.Println("Index item", Term, " does not exist :: creating new item")
		m := make(map[int]int, 0)
		m[Document] = 1
		InvertedIndexEntry := &InvertedIndexEntry{
			Term:            Term,
			Frequencies:       m,
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

// GenerateInvertedIndex for each document list
// gets each word as a token, processes it and
// generates a hash map for each document
// using them it then generates the
// inverted index of all words
func GenerateInvertedIndex(fpMaps []map[uint64]bool) InvertedIndex {

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

func Find(index InvertedIndex, Term uint64) ([]int, []int) {
	if index.HashMap[Term] != nil {
		itemPosition := index.FindItem(Term)
		item := index.Items[itemPosition]
		freqs := make([]int, len(item.DocumentListing))
		for i, d := range(item.DocumentListing) {
			freqs[i] = item.Frequencies[d]
		}
		return item.DocumentListing, freqs
	} else {
		return []int{}, []int{}
	}
}