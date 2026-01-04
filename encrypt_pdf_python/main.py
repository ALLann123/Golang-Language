#!/usr/bin/python3
import PyPDF2 as password 

out=password.PdfWriter()   #for performing actions on the pdf
original_file=password.PdfReader(open("bot_masters.pdf", "rb"))

for i in range(0, len(original_file.pages)): #iterate for number of total pages.
    out.add_page(original_file.pages[i])

newfile=open("bot_master_protected.pdf", "wb") #output file

out.encrypt("C0nque5t", use_128bit=True) #provide password
out.write(newfile)
newfile.close() #close open file

