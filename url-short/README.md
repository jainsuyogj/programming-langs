# URL shortner

Requirements:
    1) Given a URL, shorten the URL
    2) A shortened URL should not map to a single URL.
    3) Response time should be super - fast :)

Non Fucntional requirements:
    1) Always generate a URL.
    2) GET should be very fast. Create is OK.
    3) resiliency

Design:
    1) 
    6) Number of services?
    3) Load balancing? why -> there is limit on number of request which can be handled by a single server (open sockets/threads)
    4) DataBase? (key- value pair?)
    5) Unique string generation algorithm ? Generation

Pipeline, continuos build?

End-to-end testing?

error handling?
