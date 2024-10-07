# go-diamond-sequence-hints


- a annotation to sequence extractor for passing the asligned regions as hints. 
- convert your reference aligned sequence fasta to the linearize form 
- it extracts and write a fasta so that you can use the aligned regions either for the annotations or simply for generating the tags for the specific regions. 
- this is included in the go-mapper-diamond but if you have the reads to protein alignments then you can use it separately. 
```
awk '/^>/ {printf("\n%s\n",$0);next; } { printf("%s",$0);}  \
                         END {printf("\n");}' inputfasta > output.fasta
```

Gaurav Sablok
