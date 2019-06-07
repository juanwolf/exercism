package strand

// Nucleotide represents a
type Nucleotide string

const (
	adenine  Nucleotide = "A"
	cytosine Nucleotide = "C"
	guanine  Nucleotide = "G"
	thymine  Nucleotide = "T"
	uracil   Nucleotide = "U"
)

// DNAtoRNA is a map to convert DNA nucleotide to RNAs
var DNAtoRNA = map[Nucleotide]Nucleotide{
	adenine:  uracil,
	guanine:  cytosine,
	cytosine: guanine,
	thymine:  adenine,
}

// ToRNA returns the RNA equivalent of the DNA nucleotide
func (n Nucleotide) ToRNA() Nucleotide {
	return DNAtoRNA[n]
}

// ToRNA returns the rna equivalent of the DNA passed as parameters
func ToRNA(dna string) string {
	rna := ""
	for _, char := range dna {
		nucleotide := Nucleotide(char)
		rna += string(nucleotide.ToRNA())
	}
	return rna
}
