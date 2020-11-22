# -Injection
An attacker can use entry fields to inject grammatically valid constructions that subvert application 
logic. This is a way to inject malicious code into an application, most often based on SQL, NoSQL, OS, or LDAP.

Positive Technologies analysts have identified SQL injection as the most popular attack on web applications (27%).

Most commonly, injection attacks occur when a web application fails to restrict user input in site forms and entry fields.
Malicious code introduced to the application in such fields can yield access to sensitive information or even administrator rights.
![](https://user-images.githubusercontent.com/61666791/99899300-c1bfd780-2cda-11eb-90f7-065307fa0f26.png)


# -Broken Authentication
Obtaining information about the accounts of other users, or authenticating as another user, is an obvious but all-toо-common issue.

In these cases, an attacker can perform a brute-force attack. Many sites fail to enforce password complexity requirements. 
Particularly when the number of password attempts is not limited (an account is not frozen after three unsuccessful attempts, for example), 
a number of attack vectors become feasible.
![](https://user-images.githubusercontent.com/61666791/99900555-a2c54380-2ce2-11eb-8320-8f9f2f03623e.png)


# -XML External Entities (XXE)
By targeting the XML parser, an attacker can inject external entities into a document that lead to reading of local files and, in some cases, even execution of arbitrary code.

This attack is possible if the application uses XML for transmitting data between the server and user's web browser. The XML specification is potentially dangerous and can be abused to access any server files that are available to the application itself.
![](https://user-images.githubusercontent.com/61666791/99900669-aefdd080-2ce3-11eb-8fae-040678840cbe.png)
