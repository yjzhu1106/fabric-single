# node invoke.js dc AddURL D100010001 https://test.iot.org/voice00001.mp3



# node invoke.js dc GetURL D100010001



node invoke.js apmc AddPolicy '{"AS":{"userId":"user1","role":"owner","PKuser":"MAKNFKSLFHSAHFLSFA="},"AO":{"dataId":"data1","signer":"user1","sign_data":"dksalhfklhfkajlfkql","resourceKey":"be9ecaa87baefef74b9c574092466ce332513cddf981fcba08087a1be9adc5ba","url":"https://localhost"},"AP":{"auth_sign":"user1Sign","p_data":1,"p_url":1},"AE":{"createdTime":1575468182,"endTime":1976468182,"address":"10.10.39.70","sign_PKuser":"ahdjksahjkfhsjkah="}}'
node invoke.js apmc QueryPolicy ae133e4ee660cbc96a239b86e12b3367ffaea0139bb875887e46263833bd9a9e
node invoke.js apmc UpdatePolicy '{"AS":{"userId":"user1","role":"owner","PKuser":"MAKNFKSLFHSAHFLSFA="},"AO":{"dataId":"data1","signer":"user1","sign_data":"dksalhfklhfkajlfkql","resourceKey":"resource1","url":"https://localhost"},"AP":{"auth_sign":"user1Sign","p_data":1,"p_url":1},"AE":{"createdTime":1575468182,"endTime":1976468182,"address":"10.10.45.66","sign_PKuser":"ahdjksahjkfhsjkah="}}'
node invoke.js apmc QueryPolicy ae133e4ee660cbc96a239b86e12b3367ffaea0139bb875887e46263833bd9a9e
node invoke.js apmc DeletePolicy ae133e4ee660cbc96a239b86e12b3367ffaea0139bb875887e46263833bd9a9e



node invoke.js dacc RequestAccess '{"AS":{"userId":"request1","role":"request","PKuser":"MAKNFKSLFHSAHFLSFA="},"AO":{"dataId":"data1","signer":"user1","sign_data":"dksalhfklhfkajlfkql","resourceKey":"resource1","url":"https://localhost"}}'
node invoke.js apmc AddPolicy '{"AS":{"userId":"request1","role":"request","PKuser":"MAKNFKSLFHSAHFLSFA="},"AO":{"dataId":"data1","signer":"user1","sign_data":"dksalhfklhfkajlfkql","resourceKey":"resource1","url":"https://localhost"},"AP":{"auth_sign":"user1Sign","p_data":1,"p_url":1},"AE":{"createdTime":1575468182,"endTime":1976468182,"address":"10.10.39.70","sign_PKuser":"ahdjksahjkfhsjkah="}}'
node invoke.js dacc ResponseAccess '{"policyId":"eb4820c786a676b0f5485048a590aa5a29b8d52cf8181c56a3a8b7b4a894b94b","owner":"user1","requestId":"request1","status":1,"endTime":19239932,"timestamp":132173718}'
node invoke.js dacc CheckAccess eb4820c786a676b0f5485048a590aa5a29b8d52cf8181c56a3a8b7b4a894b94b
 


# node invoke.js pdpc SignData '{"data":"183cm80kg","random":128,"key":"-----BEGIN ecc private key-----\nMIHcAgEBBEIB5Rsn6WarZZ0FGo4/2prJYNYdxMvyoBLqaf/GinelALg6WFutM7si\nbPzu3zeT6FLqUcp95NoUJ0B4Is+ArT6scqWgBwYFK4EEACOhgYkDgYYABAA/1/Yb\nq1raIn+fk3qUZXyzH8oNDybYs8qPr+5BtWK3/ljC6lKv2HGmz4jLcmBATRHo8QbL\nwxFx2nvX3j+TT0LnZgDuCWZmkLlCocIJsPgUFjUf/qp+npG8NSJS2cuYS2owg3lz\ntrV7zCKtvR07e1w5mb1njItA1dDbKeIygJHMrsnTxg==\n-----END ecc private key-----\n"}' 
# node invoke.js pc AddPolicy '{"AS":{"userId":"13800010002","role":"u1","group":"g1"},"AO":{"deviceId":"D100010001","MAC":"00:11:22:33:44:55"},"AP":1,"AE":{"createdTime":1575468182,"endTime":1576468182,"allowedIP":"*.*.*.*"}}'
# node invoke.js pc QueryPolicy da1f08aa8ee250a65057c354b6952e2baf6ae3c9505cb9ca3e5ab4138e56b9d1




node invoke.js pdcc AddData '{"resourceId":"R0002","data":"183/80","url":"https://localhost","timestamp":4214121}'
node invoke.js apmc AddPolicy '{"AS":{"userId":"user001","role":"owner","PKuser":"MAKNFKSLFHSAHFLSFA="},"AO":{"dataId":"data1","signer":"user1","sign_data":"dksalhfklhfkajlfkql","resourceKey":"be9ecaa87baefef74b9c574092466ce332513cddf981fcba08087a1be9adc5ba","url":"https://localhost"},"AP":{"auth_sign":"user1Sign","p_data":1,"p_url":1},"AE":{"createdTime":1575468182,"endTime":1976468182,"address":"10.10.39.70","sign_PKuser":"MAKNFKSLFHSAHFLSFA="}}'
node invoke.js pdcc GetData ae133e4ee660cbc96a239b86e12b3367ffaea0139bb875887e46263833bd9a9e
node invoke.js pdcc GetURL ae133e4ee660cbc96a239b86e12b3367ffaea0139bb875887e46263833bd9a9e




node invoke.js apmc AddPolicy '{"AS":{"userId":"user12","role":"owner","PKuser":"MAKNFK4242432432423SLFHSAHFLSFA="},"AO":{"dataId":"data1","signer":"user1","sign_data":"dksalhf4234242342423432423klhfkajlfkql","resourceKey":"be9ecaa87baefef74b9c574092466ce332513cddf981fcba08087a1be9adc5ba","url":"https://localhost"},"AP":{"auth_sign":"user1Sign","p_data":1,"p_url":1},"AE":{"createdTime":1575468182,"endTime":1976468182,"address":"10.10.39.70","sign_PKuser":"ahdjksahjkfhsjkah="}}'


# node invoke.js pc AddPolicy '{"AS":{"userId":"13800010001","role":"u1","group":"g1"},"AO":{"deviceId":"D100010001","MAC":"00:11:22:33:44:55"},"AP":1,"AE":{"createdTime":1575468182,"endTime":1576468182,"allowedIP":"*.*.*.*"}}'
# node invoke.js pc QueryPolicy 40db810e4ccb4cc1f3d5bc5803fb61e863cf05ea7fc2f63165599ef53adf5623
# node invoke.js ac CheckAccess '{"AS":{"userId":"13800010002","role":"u1","group":"g1"},"AO":{"deviceId":"D100010001","MAC":"00:11:22:33:44:55"}}'
# node invoke.js ac CheckAccess '{"AS":{"userId":"13800010001","role":"u1","group":"g1"},"AO":{"deviceId":"D100010001","MAC":"00:11:22:33:44:55"}}'

node invoke.js pdcc AddData '{"resourceId":"user1_R0003","data":"183/80","url":"https://localhost","timestamp":4214121}' user1 "-----BEGIN ecc public key-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQA23ugWCPhBWyhScIm9O6Yvbr/m3a4\nhgAo6Rw2dVv+mh06qMu3+k+S7Azt0QwLLWBOoby1nyzde7/aNzsdlp52GWoBugoP\n3/1tZSbG0wm0rlI9qkLfDtHN0t5sVJMjFtjgPzFYA5XKrT/OW9MdcqzwEsxtoVHa\nrJg3xLOU23u3vKivUUA=\n-----END ecc public key-----\n"
node invoke.js apmc AddPolicy '{"AS":{"userId":"user1","role":"owner","PKuser":"-----BEGIN ecc public key-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQA23ugWCPhBWyhScIm9O6Yvbr/m3a4\nhgAo6Rw2dVv+mh06qMu3+k+S7Azt0QwLLWBOoby1nyzde7/aNzsdlp52GWoBugoP\n3/1tZSbG0wm0rlI9qkLfDtHN0t5sVJMjFtjgPzFYA5XKrT/OW9MdcqzwEsxtoVHa\nrJg3xLOU23u3vKivUUA=\n-----END ecc public key-----\n"},"AO":{"dataId":"data1","signer":"user1","sign_data":"dasbkfhakfh","resourceKey":"d17cebb054339d715d583a8de5683e656784ebee2399ce19f8779dab999e6b69","url":"https://localhost"},"AP":{"auth_sign":{"rtext":"1959251352653671459012865375812608926140316549573057863114886951502070552298855612060536111663701109539128016721844094106655494517403423971007724277207957280","stext":"1713192365449443888414985415775159560806012411684699009918619676790799078402271467630730094231606254861296090718373543777042893800688279078328777432209067612"},"p_data":1,"p_url":1},"AE":{"createdTime":1575468182,"endTime":1976468182,"address":"10.10.39.70","sign_PKuser":{"rtext":"415985477922990628170973440735556282936584277728487414981350898378585420491886683633796104335214673404579507405052235690231828976329654425431047086712903632","stext":"413766465868270820610562025056694113779883052278244010208294697352274489465935593872488429524967847498016545975992179469545326807881814174952753699657084418"}}}'
node invoke.js pdcc GetData 5efce58af1d20cf431693efce3ccfdaa8027b96039492c3431aff5881ec17b9b

node invoke.js dacc RequestAccess '{"AS":{"userId":"user2","role":"request","PKuser":"-----BEGIN ecc public key-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQA23ugWCPhBWyhScIm9O6Yvbr/m3a4\nhgAo6Rw2dVv+mh06qMu3+k+S7Azt0QwLLWBOoby1nyzde7/aNzsdlp52GWoBugoP\n3/1tZSbG0wm0rlI9qkLfDtHN0t5sVJMjFtjgPzFYA5XKrT/OW9MdcqzwEsxtoVHa\nrJg3xLOU23u3vKivUUA=\n-----END ecc public key-----\n"},"AO":{"dataId":"data1","signer":"user1","sign_data":"nil","resourceKey":"d17cebb054339d715d583a8de5683e656784ebee2399ce19f8779dab999e6b69","url":"https://localhost"}}'
node invoke.js apmc AddPolicy '{"AS":{"userId":"user2","role":"request","PKuser":"-----BEGIN ecc public key-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQA23ugWCPhBWyhScIm9O6Yvbr/m3a4\nhgAo6Rw2dVv+mh06qMu3+k+S7Azt0QwLLWBOoby1nyzde7/aNzsdlp52GWoBugoP\n3/1tZSbG0wm0rlI9qkLfDtHN0t5sVJMjFtjgPzFYA5XKrT/OW9MdcqzwEsxtoVHa\nrJg3xLOU23u3vKivUUA=\n-----END ecc public key-----\n"},"AO":{"dataId":"data1","signer":"user1","sign_data":"dasbkfhakfh","resourceKey":"d17cebb054339d715d583a8de5683e656784ebee2399ce19f8779dab999e6b69","url":"https://localhost"},"AP":{"auth_sign":{"rtext":"1959251352653671459012865375812608926140316549573057863114886951502070552298855612060536111663701109539128016721844094106655494517403423971007724277207957280","stext":"1713192365449443888414985415775159560806012411684699009918619676790799078402271467630730094231606254861296090718373543777042893800688279078328777432209067612"},"p_data":1,"p_url":1},"AE":{"createdTime":1575468182,"endTime":1976468182,"address":"10.10.39.70","sign_PKuser":{"rtext":"415985477922990628170973440735556282936584277728487414981350898378585420491886683633796104335214673404579507405052235690231828976329654425431047086712903632","stext":"413766465868270820610562025056694113779883052278244010208294697352274489465935593872488429524967847498016545975992179469545326807881814174952753699657084418"}}}'
node invoke.js dacc ResponseAccess '{"policyId":"9972ca74aa209bf80095df7bffb88198cd00e1dd9c5680a52913ee5ef9d3b728","owner":"user1","requestId":"user2","status":1,"endTime":19239932,"timestamp":132173718}'
node invoke.js dacc CheckAccess 9972ca74aa209bf80095df7bffb88198cd00e1dd9c5680a52913ee5ef9d3b728
node invoke.js dacc GetRequest 66591886e057e11b71cfc3eecd8dffbe7f70d06797f96de4d39a452274600084
node invoke.js dacc GetResponse 42feb6bac564d45a61e25dcfc61aa8337efaee2c9a89a77bc120f3f65b4fd011
node invoke.js vlrc GetRecord user1_R0001
node invoke.js apmc DeletePolicy 9972ca74aa209bf80095df7bffb88198cd00e1dd9c5680a52913ee5ef9d3b728
node invoke.js apmc QueryPolicy 9972ca74aa209bf80095df7bffb88198cd00e1dd9c5680a52913ee5ef9d3b728