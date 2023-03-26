.PHONY: lab01_zip
lab01_zip:
	mkdir Rys_Szymon_1
	mkdir -p archives
	cp -r lab01/ Rys_Szymon_1/
	zip -r archives/Rys_Szymon_1.zip Rys_Szymon_1/ oswiadczenie.txt
	rm -r Rys_Szymon_1

.PHONY: lab02_zip
lab02_zip:
	mkdir Rys_Szymon_2
	mkdir -p archives
	cp -r lab02/ Rys_Szymon_2/
	zip -r archives/Rys_Szymon_2.zip Rys_Szymon_2/ oswiadczenie.txt
	rm -r Rys_Szymon_2
