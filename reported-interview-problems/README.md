# Reported Yandex Go interview problems

These are reconstructions of problems reported by candidates or community interview banks. They are not official Yandex problem statements and may omit clarifications that an interviewer supplied verbally.

Retrieved and reconstructed on August 24, 2026. Statements are paraphrased; source links are preserved for provenance. Generated answers from source sites are intentionally not copied because several omit safe cancellation, joining, or channel-lifecycle handling.

| ID | Problem | Concurrency focus | Evidence |
|---|---|---|---|
| YR01 | [Kafka-to-ClickHouse acknowledged pipeline](yr01-kafka-clickhouse-pipeline/README.md) | Pipeline stages, batching, ordered commits, failure propagation | High: current interview recollection matches a Yandex-tagged bank entry |
| YR02 | [Least-loaded client-side balancer](yr02-least-loaded-balancer/README.md) | Concurrent load accounting, mutex scope, cleanup | Medium: Yandex-tagged community bank entry |
| YR03 | [Resumable large-table copier](yr03-resumable-table-copy/README.md) | Checkpointing, idempotency; bounded parallel extension | Medium: Yandex-tagged community bank entry |
| YR04 | [Bounded concurrent sitemap crawler](yr04-sitemap-crawler/README.md) | Dynamic work, deduplication, bounded workers, completion detection | High: detailed first-person 2026 interview report |

## Sources

- [YR01 community report](https://sobes.tech/en/bank/go/f5db5534-99f2-4d3e-b26f-8481e1414783)
- [YR01 alternate duplicate](https://sobes.tech/en/bank/go/b1078edb-700a-4fc6-a489-5050e83fc792)
- [YR02 community report](https://sobes.tech/en/bank/go/d3c5335d-aedc-4b3d-b4e6-a67d3902c241)
- [YR03 bank page containing the report](https://sobes.tech/en/bank/go?page=232)
- [YR04 first-person Habr report](https://habr.com/ru/articles/1006022/)
- [Official Yandex backend interview overview](https://yandex.ru/jobs/interview/backend)

## How to use these

Treat every ambiguity as part of the exercise. Before coding, ask about termination, cancellation, method concurrency safety, ordering, error precedence, retry semantics, and ownership of returned slices. That is especially important for YR01: its recovered interface has no context and does not define the end-of-stream signal.
