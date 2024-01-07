# Chat Server
A simple Server, for sending and loading unencrypted messages

## Routes
| Path      | Form Values                                                                                     | Queries                                  | Returns                           |
|-----------|-------------------------------------------------------------------------------------------------|------------------------------------------|-----------------------------------|
| /register | username: The alias of the user<br>password: used for logging in and generating the users token | -                                        | -                                 |
| /login    | username: The alias of the user<br>password: the password, set when registering                 | -                                        | Users token                       |
| /send     | token: The token, obtained when logged in<br>message: the message, which is wanted to send      | chat: the chats id, obtained with /chats | -                                 |
| /load     | -                                                                                               | id: the chats id                         | All the chats messages, formatted |
| /chats    | -                                                                                               | -                                        | All Chats, ID and Name            |
