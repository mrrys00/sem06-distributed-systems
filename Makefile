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

.PHONY: lab03_zip
lab03_zip:
	mkdir Rys_Szymon_3
	mkdir -p archives
	cp -r lab03/ Rys_Szymon_3/
	zip -r archives/Rys_Szymon_3.zip Rys_Szymon_3/ oswiadczenie.txt
	rm -r Rys_Szymon_3

.PHONY: lab04_zip
lab04_zip:
	mkdir Rys_Szymon_4
	mkdir -p archives
	cp -r lab04/ Rys_Szymon_4/
	zip -r archives/Rys_Szymon_4.zip Rys_Szymon_4/ oswiadczenie.txt
	rm -r Rys_Szymon_4
