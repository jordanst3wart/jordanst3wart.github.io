---
title: Writing a terraform variables parser and writer
date: '2025-02-25'
tags: ['technology', 'parser', 'hcl']
summary: A short note about writing a terraform parser for tfvar files.
---

I recently wrote a small library to parse tfvar files, or hcl files, for jdk languages. I wrote the library in kotlin, and it's just two files, with a lot of tests. I wrote it so, because I wanted to be a library author of a really small but useful library. Currently, there is only hcl4j that can parse hcl files, but not write HCL files. I really wanted it to be simple, and low/no dependencies.

I found the hardest part of about writing a library was actually publishing the package to maven. Writing the parser was challenging too, but I think sorting out PGP keys was more challenging. I wanted the library to be maintained in a test driven way, which I think makes sense with this library. It's easy to have a lot of tests that parse HCL, and match an expected result.

I want to compare it to hcl4j...
Number of tests:
hcl4j: 52 tests
hcl: 31 tests

Number of lines of code:
hcl: 439
hcl4j: 3578 (also generates 1678 lines with jflex)

Number of source code files:
hcl: 2
hcl4j: 36 (plus one generated)

I wonder how they got heredoc parsing to work.

Indented heredoc I still don't have support for, but this is normal heredoc:

```hcl
user_data2 = <<FOO
#!/bin/bash
echo "Hello, World!"
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd
FOO
```

It has a delimiter, which can be any word repeated, in this case 'FOO', and starts with `<<`. It starts with `<<-` for indented heredoc.

Hcl4j parses heredoc fine, but doesn't get the indent right with indented heredoc.

ie.

```hcl
user_data2 = <<-FOO
  #!/bin/bash
  echo "Hello, World!"
  yum update -y
  yum install -y httpd
  systemctl start httpd
  systemctl enable httpd
FOO
```

Should be processed to:

```bash
#!/bin/bash
echo "Hello, World!"
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd
```

The two spaces should be removed.
