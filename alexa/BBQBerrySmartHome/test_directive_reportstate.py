import json
import handler


request = {
    "directive": {
        "endpoint": {
            "cookie": {},
            "endpointId": "bbqberry",
            "scope": {
                "token": "Atza|IwEBIIM1Jk0zzEZNvrskXb9ofwhSfx_zO5JPcWvJ6_WLN93ZIHFK9sGDF2xTgU8UxDLmkaWGjL843ggSixnpbbhokriKbCR1rcA1USxWBOvI5BM3THWg7swp4ySawTWJ0EOLsRhIhnj2-KN4xCxFpMKF4FA97yuHLz0wMdieaI1X5cr3y5ASUP4q1qcSXWFCe753cRQOiJPyDdIxW_PKRbzD5pfsIvZ5IVrfwx5qmdo1PsA7cKqSjHq9Gvc9eg93uklHyuOR7s9XJeXo6r1Wn0UMhbQmI4Vl428ve8jbIKZwD16M_N-yUaN9NTBEw3mK2gak5KjvnHp8ENFvtS5s9Yy_Lvw5qd4zKctTptzBTD-RuQuytwB73RxJ8uA4Eqdspf1In30J8rU7Ty795CS8-dhD8Uzaw6eTHdZeIQMPGVGf7J7Z--PIXKENWHjcP4EwTs2npyEJuMXCNqAbpVylC5clgNKxYTJuaF7Gmdh3SSldxc9bVxyqZ8kX_xKXr4Pkwm5Sbwc",
                "type": "BearerToken"
            }
        },
        "header": {
            "correlationToken": "AAAAAAAAAADBZt5qTLHppE/m3LrCcUjhBAIAAAAAAAB7MwFlwRP+4Q7XjVaiZ/oznNcsABhUAFG90MpB2k9gHGpKjBVmwMR9wHhaBV79kAeYUN2QTyi6Fy+u41xl/3JT3Md5oPHCL17PMrhoN8PlzvAz78aLUfvdV7VjFjrL9wM/T3MRD/s2KmcGb8MDVAHByNn9HOEo6TouA18b1Yo+WiMxqHcIYlrvnmpSBjJLO+bg4SzanT4gzjB9iAuf1BEPa6Xott8fI/iZ5XzQq2B6B/83wH2ez4Y5AwpotlK4/Z5ALew+QS2oL69RtWha7rfzF2wsvDWbsSiU8KAV6ihHT9tP4Ug0K2/kToYN6PtdYOvk7TUtwV4QitupCxuj4dY0xS96dvXgprWQ5w7QU+SlQxiZhFMZx6G2JgqLtBig7Fr9s8QO1PYFumhGrNLshPh627AHJX/yuJY8md4pjznZAdcPwK7ko94qG1GCbipjewzoaQYXVz8+5dK1jGAmjT4lV3Osc5n/olvySuMMG83hY0wbvzx5jf0lBdrG1jorKIYVmxlbattfszLh0AOddORbOd0pAyk//pU01kSCV38Yw244bzEWCr7+SMqiVanAAQN7mAqE3kxW+mfJTwo0zhttOaKTTUMTUbFbLl0Wq4P7rrYtWM2BBC91zG18o1mz6DEY/Z9p/0La15IK+rhxPpxbxCf0YOOAKJe/DAcAODY3MGZxy/E=",
            "messageId": "bb819a53-791f-48a6-8089-c2ef5fafa46b",
            "name": "ReportState",
            "namespace": "Alexa",
            "payloadVersion": "3"
        },
        "payload": {}
    }
}

response = handler.lambda_handler(request, None)
print(json.dumps(response, indent=4))
