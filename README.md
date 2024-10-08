# go-diamond-sequence-hints


- a annotation to sequence extractor for passing the asligned regions as hints. 
- convert your reference aligned sequence fasta to the linearize form 
- it extracts and write a fasta so that you can use the aligned regions either for the annotations or simply for generating the tags for the specific regions. 
- this is included in the go-mapper-diamond but if you have the reads to protein alignments then you can use it separately. 
```
awk '/^>/ {printf("\n%s\n",$0);next; } { printf("%s",$0);}  \
                         END {printf("\n");}' inputfasta > output.fasta
```

```
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-diamond-sequence-hints ±main⚡ » \
go run main.go -h
Analyzer for the diamond aligner and pacbio reads for hints

Usage:
  get [command]

Available Commands:
  UpStream
  alignment
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for get

Use "get [command] --help" for more information about a command.
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-diamond-sequence-hints ±main⚡ » \
go run main.go alignment -h
Analyzes the hsp from the diamond read to protein alignment

Usage:
  get alignment annotator [flags]

Flags:
  -a, --alignmentfile string    alignment (default "alignment file to be analyzed")
  -h, --help                    help for alignment
  -p, --referencefasta string   pacbio file (default "pacbio reads file")
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-diamond-sequence-hints ±main⚡ » \
go run main.go Upstream -h
specific for the genome alignment regions upstream and the downstream of the alignments

Usage:
  get Upstream Aligner [flags]

Flags:
  -a, --alignmentfile string             alignment (default "alignment file to be analyzed")
  -d, --downstream of the hsp tags int   downstream tags (default 5)
  -h, --help                             help for Upstream
  -p, --referencefasta string            pacbio file (default "pacbio reads file")
  -u, --upstream of the hsp tags int     upstream tags (default 4)
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-diamond-sequence-hints ±main⚡ » \
go run main.go UpStream -a matches.tsv -p pacbioreads.fasta -u 10 -d 10                            1 ↵
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-diamond-sequence-hints ±main⚡ » \
ls -la
total 88
drwxr-xr-x. 1 gauavsablok gauavsablok   256 Oct  8 11:33 .
drwxr-xr-x. 1 gauavsablok gauavsablok  1670 Oct  7 21:05 ..
drwxr-xr-x. 1 gauavsablok gauavsablok   148 Oct  8 08:00 .git
-rw-r--r--. 1 gauavsablok gauavsablok   211 Oct  7 23:51 go.mod
-rw-r--r--. 1 gauavsablok gauavsablok   896 Oct  7 23:51 go.sum
-rw-r--r--. 1 gauavsablok gauavsablok  5810 Oct  8 11:33 main.go
-rwxr-xr-x. 1 gauavsablok gauavsablok   324 Oct  8 11:08 matches.tsv
-rw-r--r--. 1 gauavsablok gauavsablok 54117 Oct  7 20:59 pacbioreads.fasta
-rw-r--r--. 1 gauavsablok gauavsablok   624 Oct  8 00:04 README.md
-rw-r--r--. 1 gauavsablok gauavsablok   304 Oct  8 11:32 sequences-annotation.txt
-rw-r--r--. 1 gauavsablok gauavsablok   384 Oct  8 11:33 sequences-annotation-upstream-downstream.txt

```

Gaurav Sablok
