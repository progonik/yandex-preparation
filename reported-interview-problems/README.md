# Yandex-style Go live-coding problems

These are candidate-facing reconstructions of problems reported by Yandex interview candidates. They are not official Yandex statements. Each task is written so that an interviewer can reveal the concurrency follow-up after the basic version works.

| ID | Problem | Main topic | Suggested time |
|---|---|---|---|
| YR01 | [Acknowledged batch pipeline](yr01-kafka-clickhouse-pipeline/README.md) | Batching, pipeline stages, ordered commits | 60 minutes |
| YR02 | [Least-loaded client-side balancer](yr02-least-loaded-balancer/README.md) | Shared state, load accounting, mutex scope | 45–60 minutes |
| YR03 | [Resumable large-table copy](yr03-resumable-table-copy/README.md) | Checkpoints, bounded batches, recovery | 60 minutes |
| YR04 | [Concurrent sitemap crawler](yr04-sitemap-crawler/README.md) | Dynamic work, deduplication, bounded concurrency | 60 minutes |

## Interview mode

For each problem:

1. Read only the context and the basic task.
2. Spend up to five minutes asking clarifying questions.
3. State the important invariants before writing code.
4. Implement and test the basic version.
5. Continue to the concurrency follow-up when the basic version is correct.

The questions at the end of each statement are intentionally unanswered. In a mock interview, the interviewer should choose an answer and make the candidate repeat the resulting requirement before coding.

## Provenance

The statements were reconstructed on August 24, 2026 and paraphrased from these reports:

- [YR01 community report](https://sobes.tech/en/bank/go/f5db5534-99f2-4d3e-b26f-8481e1414783)
- [YR01 duplicate report](https://sobes.tech/en/bank/go/b1078edb-700a-4fc6-a489-5050e83fc792)
- [YR02 community report](https://sobes.tech/en/bank/go/d3c5335d-aedc-4b3d-b4e6-a67d3902c241)
- [YR03 report listing](https://sobes.tech/en/bank/go?page=232)
- [YR04 first-person interview report](https://habr.com/ru/articles/1006022/)
- [Official Yandex backend interview overview](https://yandex.ru/jobs/interview/backend)
