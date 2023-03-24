 # Go Service Evidence Client
 
Is a client library to make working with [Service Evidence](https://github.com/westwardharbor0/server-evidence) easy.

# Example usage

Simple example usage:
 ```go

 // Set up the client with url to service.
 client := ServerEvidenceClient{
     Host: "http://localhost:8080",
 }
 // Get all the machines.
 machines, err := client.All()
 fmt.Println(machines)

 // Update one machine attributes and save it to service.
 machines[0].IPV4 = "1.2.3.4"
 err := client.Update(machines[0])

 ```
