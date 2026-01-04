#!/usr/bin/python3
import PyPDF2 as password 

def addPassword(fileLocation, passwordSet, newFileName):
    out=password.PdfWriter()   #for performing actions on the pdf
    original_file=password.PdfReader(open(fileLocation, "rb"))

    for i in range(0, len(original_file.pages)): #iterate for number of total pages.
        out.add_page(original_file.pages[i])
    newfile=open(newFileName, "wb") #output file

    out.encrypt(passwordSet, use_128bit=True) #provide password
    out.write(newfile)
    newfile.close() #close open file

print("***************"*3)
print("        PDF PASSWORD SETTER     ")
print("***************"*3)

targetFile= input("Enter File Location/Name: ")
desiredPass=input("Password to encrypt>> ")
print("\n")
outputfile=input("Name of encrypted file>> ")
addPassword(targetFile, desiredPass, outputfile)
print("[+]DOne.....")

"""
    cmd>> python pdf_set_pass.py
*********************************************
        PDF PASSWORD SETTER
*********************************************
Enter File Location/Name: error.pdf
Password to encrypt>> Mr.Robot2025


Name of encrypted file>> locked_err.pdf
[+]DOne.....

"""