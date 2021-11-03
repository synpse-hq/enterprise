build-ssh-keygen:
	go build -o bin/sshkeygen ./hacks/ssh-key/ 

generate-ssh-key:
	./bin/sshkeygen