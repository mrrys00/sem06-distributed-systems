.PHONY: lab01_zip
lab01_zip:
	mkdir Rys_Szymon_1
	mkdir -p archives
	re="(~*.pdf)/"
	cp -r lab01/$(re) Rys_Szymon_1/
	zip -r archives/Rys_Szymon_1.zip Rys_Szymon_1/ oswiadczenie.txt
	rm -r Rys_Szymon_1
