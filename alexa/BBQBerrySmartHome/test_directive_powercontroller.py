import json
import handler


request = {
    "directive": {
        "endpoint": {
            "cookie": {},
            "endpointId": "endpoint-002",
            "scope": {
                "token": "Atza|IwEBIIM1Jk0zzEZNvrskXb9ofwhSfx_zO5JPcWvJ6_WLN93ZIHFK9sGDF2xTgU8UxDLmkaWGjL843ggSixnpbbhokriKbCR1rcA1USxWBOvI5BM3THWg7swp4ySawTWJ0EOLsRhIhnj2-KN4xCxFpMKF4FA97yuHLz0wMdieaI1X5cr3y5ASUP4q1qcSXWFCe753cRQOiJPyDdIxW_PKRbzD5pfsIvZ5IVrfwx5qmdo1PsA7cKqSjHq9Gvc9eg93uklHyuOR7s9XJeXo6r1Wn0UMhbQmI4Vl428ve8jbIKZwD16M_N-yUaN9NTBEw3mK2gak5KjvnHp8ENFvtS5s9Yy_Lvw5qd4zKctTptzBTD-RuQuytwB73RxJ8uA4Eqdspf1In30J8rU7Ty795CS8-dhD8Uzaw6eTHdZeIQMPGVGf7J7Z--PIXKENWHjcP4EwTs2npyEJuMXCNqAbpVylC5clgNKxYTJuaF7Gmdh3SSldxc9bVxyqZ8kX_xKXr4Pkwm5Sbwc",
                "type": "BearerToken"
            }
        },
        "header": {
            "correlationToken": "AAAAAAAAAADBZt5qTLHppE/m3LrCcUjhBAIAAAAAAABQi6+Tqc1qP5jYTnHYvPTb+pTAueq+0QJEYYiyHrWHp6or7jkLkiKgXjIMkzyi0mXr0uabsTMB/YgF9kJfwbTq+d5fF3iTL3PCiOqrT8ZDCgejTwMbRLgO+yJY+Z62Xz6KN5jV49rBCAkJtpHoWPQ4JHVlBQ+btgIi1pNXFKG1eHeDoDUO9c1tsmnfL+rPpH02Ips2TsVbjF78qy9gpnd2/l4Hj/4cavCmMY/iYVu5Dao/sYLJbcHUk/QooSP3aw08bZV8QHuQkUQdEFOE6B440IDofZzBQ0zXURLrnjoJXXKdYxC6dpSkO+rcaB1R9jKypXXgJQaMrKNgOZUO/2ogZJNqnrJNBgGue6HAVXcsOPeP3felRjwKo/M98xZTfsC1NBrqL/GsaDMSdqBLtmMw7ZB7BBRluDqJ799THAbXsM954VE9xDFO+3OZx9x0IVMh+x+9QBc3w+MzI1NMuMdSXWWJJUFOa1ks0cpzgJZfAsSX/VA/+tTLFspe/+1O0yqrlFWbGG+Z/bMMKw0t3Ag/Q8xO537aLTnV03Gheof9Kbh3zc3TO/kpulEpNOXqGvrMskjBBEQCOni9m0tw/Ij+hsVxShacm7r7Up4AOtnn9VIzvlw8b+hSO0K4qsejvTfrNw6NAljEP6MGVWTBm0g757JPPuFxrA0x+DUn93//uyCvljA=",
            "messageId": "07f0bdff-c783-42f7-86b3-1a3ea92e732f",
            "name": "TurnOn",
            "namespace": "Alexa.PowerController",
            "payloadVersion": "3"
        },
        "payload": {}
    }
}

response = handler.lambda_handler(request, None)
print(json.dumps(response, indent=4))
