package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-4

Getting the aligned HSPs from the sequences so that you can use them for the annotations or tags generation.

*/

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var (
	alignmentfile  string
	referencefasta string
	upstreamStart  int
	downstreamEnd  int
)

var rootCmd = &cobra.Command{
	Use:  "get alignments",
	Long: "Analyzer for the diamond aligner and pacbio reads for hints",
}

var alignmentCmd = &cobra.Command{
	Use:  "alignment annotator",
	Long: "Analyzes the hsp from the diamond read to protein alignment",
	Run:  getSeqFunc,
}

var promoteralignmentCmd = &cobra.Command{
	Use:  "UpStream Aligner",
	Long: "specific for the genome alignment regions upstream and the downstream of the alignments",
	Run:  upstreamFunc,
}

func init() {
	alignmentCmd.Flags().
		StringVarP(&alignmentfile, "alignmentfile", "a", "alignment file to be analyzed", "alignment")
	alignmentCmd.Flags().
		StringVarP(&referencefasta, "referencefasta", "p", "pacbio reads file", "pacbio file")
	promoteralignmentCmd.Flags().
		StringVarP(&alignmentfile, "alignmentfile", "a", "alignment file to be analyzed", "alignment")
	promoteralignmentCmd.Flags().
		StringVarP(&referencefasta, "referencefasta", "p", "pacbio reads file", "pacbio file")
	promoteralignmentCmd.Flags().
		IntVarP(&upstreamStart, "upstream of the hsp tags", "u", 4, "upstream tags")
	promoteralignmentCmd.Flags().
		IntVarP(&downstreamEnd, "downstream of the hsp tags", "d", 5, "downstream tags")

	rootCmd.AddCommand(alignmentCmd)
	rootCmd.AddCommand(promoteralignmentCmd)
}

func getSeqFunc(cmd *cobra.Command, args []string) {
	refID := []string{}
	alignID := []string{}
	refIdenStart := []float64{}
	refIdenEnd := []float64{}
	alignIdenStart := []float64{}
	alignIdenEnd := []float64{}
	fOpen, err := os.Open(alignmentfile)
	if err != nil {
		log.Fatal(err)
	}

	fRead := bufio.NewScanner(fOpen)

	for fRead.Scan() {
		line := fRead.Text()
		refID = append(refID, strings.Split(string(line), "\t")[0])
		alignID = append(alignID, strings.Split(string(line), "\t")[1])
		start1, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[6], 32)
		end1, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[7], 32)
		start2, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[8], 32)
		end2, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[9], 32)
		refIdenStart = append(refIdenStart, start1)
		refIdenEnd = append(refIdenEnd, end1)
		alignIdenStart = append(alignIdenStart, start2)
		alignIdenEnd = append(alignIdenEnd, end2)
	}

	seqIDType, seqSeqType := readRef()

	type extractSeq struct {
		extractPartID  string
		extractPartSeq string
	}

	extractPartSeq := []extractSeq{}

	for i := range seqIDType {
		for j := range refID {
			if seqIDType[i] == refID[j] {
				extractPartSeq = append(extractPartSeq, extractSeq{
					extractPartID:  seqIDType[i],
					extractPartSeq: seqSeqType[i][int(refIdenStart[j]):int(refIdenEnd[j])],
				})
			}
		}
	}

	file, err := os.Create("sequences-annotation.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := range extractPartSeq {
		_, err := file.WriteString(">" + extractPartSeq[i].extractPartID + "\n" + extractPartSeq[i].extractPartSeq + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func upstreamFunc(cmd *cobra.Command, args []string) {
	refID := []string{}
	alignID := []string{}
	refIdenStart := []float64{}
	refIdenEnd := []float64{}
	alignIdenStart := []float64{}
	alignIdenEnd := []float64{}
	fOpen, err := os.Open(alignmentfile)
	if err != nil {
		log.Fatal(err)
	}

	fRead := bufio.NewScanner(fOpen)

	for fRead.Scan() {
		line := fRead.Text()
		refID = append(refID, strings.Split(string(line), "\t")[0])
		alignID = append(alignID, strings.Split(string(line), "\t")[1])
		start1, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[6], 32)
		end1, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[7], 32)
		start2, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[8], 32)
		end2, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[9], 32)
		refIdenStart = append(refIdenStart, start1)
		refIdenEnd = append(refIdenEnd, end1)
		alignIdenStart = append(alignIdenStart, start2)
		alignIdenEnd = append(alignIdenEnd, end2)
	}

	seqIDType, seqSeqType := readRef()

	type extractStreamSeq struct {
		extractStreamPartID  string
		extractStreamPartSeq string
	}

	extractupstreamPartSeq := []extractStreamSeq{}

	for i := range seqIDType {
		for j := range refID {
			if seqIDType[i] == refID[j] {
				extractupstreamPartSeq = append(extractupstreamPartSeq, extractStreamSeq{
					extractStreamPartID:  seqIDType[i],
					extractStreamPartSeq: seqSeqType[i][int(int(refIdenStart[j])-upstreamStart):int(downstreamEnd+int(refIdenEnd[j]))],
				})
			}
		}

		file, err := os.Create("sequences-annotation-upstream-downstream.txt")
		if err != nil {
			log.Fatal(err)
		}
		for i := range extractupstreamPartSeq {
			_, err := file.WriteString(">" + extractupstreamPartSeq[i].extractStreamPartID + "\n" + extractupstreamPartSeq[i].extractStreamPartSeq + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func readRef() ([]string, []string) {

	extractID := []string{}
	seqID := []string{}

	readF, err := os.Open(referencefasta)
	if err != nil {
		log.Fatal(err)
	}

	openF := bufio.NewScanner(readF)

	for openF.Scan() {
		line := openF.Text()
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" ||
			string(line[0]) == "C" {
			seqID = append(seqID, line)
		}
		if string(line[0]) == ">" {
			extractID = append(extractID, strings.ReplaceAll(string(line), ">", ""))
		}
	}
	return extractID, seqID
}
