# parse-motion-tasks

Simple program which converts a Motion exported markdown style table into a list of tasks for llama-life

### Usage

The following text give as output:

```text

| Name | Duration|
| --- | --- |
| [Title of the task](link) | 1h |

```

```text

Title of the task 60
```

### Usage

```bash
pbpaste | parse-motion-tasks | pbcopy
```