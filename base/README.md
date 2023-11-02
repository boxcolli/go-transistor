# /base
- Base doesn't support relative complement of topic.
    - ex)
        ```
        ADD [A0]
        DEL [A0 B0 C0] - this request will be ignored.
        result: [A0]
        ```
- Base supports union of topic.
    - ex) 
        ```
        ADD [A0 B0 C0]
        ADD [A0]
        result: [A0] - B0 and C0 is truncated.
        ```
