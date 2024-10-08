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

func init() {
	alignmentCmd.Flags().
		StringVarP(&alignmentfile, "alignmentfile", "a", "alignment file to be analyzed", "alignment")
	alignmentCmd.Flags().
		StringVarP(&referencefasta, "referencefasta", "p", "pacbio reads file", "pacbio file")

	rootCmd.AddCommand(alignmentCmd)
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
			extractID = append(extractID, line)
		}
	}
	type extractSeq struct {
		id  string
		seq string
	}

	storeSeq := []extractSeq{}
	for i := range extractID {
		for j := range refIdenStart {
			if extractID[i] == refID[j] {
				storeSeq = append(storeSeq, extractSeq{
					id:  extractID[j],
					seq: seqID[j][int(refIdenStart[i]):int(refIdenEnd[i])],
				})
			}
		}
	}

	file, err := os.Create("sequencesannotation.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := range storeSeq {
		_, err := file.WriteString(">" + storeSeq[i].id + "\n" + storeSeq[i].seq)
		if err != nil {
			log.Fatal(err)
		}
	}
}
