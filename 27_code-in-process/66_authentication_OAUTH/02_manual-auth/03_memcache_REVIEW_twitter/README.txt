ENCRYPT / HASH passwords
- make your hash slow
--- prevents brute force
- salt your hash
--- store your salt with your hash so you can use it
--- each salt is unique to each password
- Corey used bcrypt to do this:
golang.org/x/crypto/bcrypt
- golang.org/x/ are maybe experimental projects made by google and not yet in standard library
- you could also us a second salt, called a pepper sometimes
-- this is a code unique to the whole website
-- it's not stored with the passwords
https://github.com/Saxleader/fall2015/tree/master/code-review_improvement
