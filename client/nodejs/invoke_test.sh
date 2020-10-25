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