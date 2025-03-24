# edot-test-case

# Greetings
Hi eDOT Team,

Thank you for the opportunity so I can take this test.

After reviewing the scope of the take-home test, I find it quite extensive to complete within five days (plus the three-day extension) given my current workload. In my professional experience, this is the first time I have encountered a take-home test with such a large MVP.

In a real-world scenario, I would typically consider one of the following approaches:
1. Reducing the scope or breaking the delivery into multiple batches.
2. Prioritizing functionality completion while deferring unit and integration testing (considering it as technical debt), which may impact the overall quality of the delivered application.
3. Extending the deadline to ensure a more complete and well-tested solution.

I sincerely appreciate the three-day deadline extension. However, given my current office workload, I may not be able to fully take advantage of the extra time to achieve the quality I would ideally aim for.

Given the circumstances, I have decided to prioritize functionality completion. This means I will focus on implementing the core features within the given timeframe while deferring unit and integration testing as technical debt. In a real-world scenario, this approach would have trade-offs, particularly regarding test coverage and overall code quality. However, it ensures that the primary functionality is in place and operational by the deadline.

Thank you for your understanding.

## Contents
This repository contains multiple microservices, each serving a distinct function. Ideally, each microservice should have its own dedicated GitHub repository to maintain separation of concerns, facilitate independent deployments, and improve scalability.

However, for the purpose of this work submission, I have consolidated all microservices into a single GitHub repository. This decision was made to simplify the review process and ensure all related components are easily accessible in one place.

In a real-world scenario, each microservice would typically be in its own repository, following best practices for microservices architecture.

## Future improvements
1. Record stock movement.
2. Use a distributed lock on CRON jobs.
3. Implement asynchronous stock release using a message broker.
4. Perform stock release using orderNo.

   Note: Stock movement recording (Point 1) must be completed first to ensure accurate tracking before implementing this.

5. Implement HMAC signature authentication for API payment confirmation.
6. Add a Redis lock for stock deduction.

## Specifications
Apps:
- edot-user-rest
    - address:
        - Docker -> :9100
        - Kubernetes -> :30100
- edot-shop-rest
    - address:
        - Docker -> :9200
        - Kubernetes -> :30200
- edot-product-rest
    - address:
        - Docker -> :9300
        - Kubernetes -> :30300
- edot-warehouse-rest
    - address:
        - Docker -> :9400
        - Kubernetes -> :30400
- edot-order-rest
    - address:
        - Docker -> :9500
        - Kubernetes -> :30500
- edot-order-cron

PostgreSQL:
- user
    - address:
        - Docker -> :5532
        - Kubernetes -> : 30600
    - username: user
    - password: password
    - database: edot_user_db
- shop
    - address:
        - Docker -> :5632
        - Kubernetes -> : 30700
    - username: user
    - password: password
    - database: edot_shop_db
- product
    - address:
        - Docker -> :5732
        - Kubernetes -> : 30800
    - username: user
    - password: password
    - database: edot_product_db
- warehouse
    - address:
        - Docker -> :5832
        - Kubernetes -> : 30900
    - username: user
    - password: password
    - database: edot_warehouse_db
- order
    - address:
        - Docker -> :5932
        - Kubernetes -> : 31000
    - username: user
    - password: password
    - database: edot_order_db
