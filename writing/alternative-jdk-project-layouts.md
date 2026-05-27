---
title: Alternative jdk project layouts
date: '2025-1-19'
tags: ['jdk', 'maven', 'bazel', 'clojure', 'leiningen', 'kotlin']
summary: Alternatives jdk project layouts
---

I saw this project layout in Clojure using the Leiningen build tool:

```
$ tree -F -a --dirsfirst my-stuff/

my-stuff/
├── doc/
│   └── intro.md
├── resources/
├── src/
│   └── my_stuff/
│       └── core.clj
├── test/
│   └── my_stuff/
│       └── core_test.clj
├── CHANGELOG.md
├── .gitignore
├── .hgignore
├── LICENSE
├── project.clj
└── README.md
```

ref: https://codeberg.org/leiningen/leiningen/src/branch/stable/doc/TUTORIAL.md#creating-a-project

It blows my mind that you can have a single folder `my_stuff` as the namespace folder. I learn the importance of jdk namespaces from Rich Hickey, the creator of clojure. He said it's really important to not have namespaces leak out into the system. So It's surprising to see this. I always assume there is some hidden danger in not using really long namespaces. But if a well used build tools recommends this approach we might be able to skip at least some of the long namespaces around. Such as you can see in enterprise buzzfizz here: https://github.com/EnterpriseQualityCoding/FizzBuzzEnterpriseEdition/tree/uinverse/src/main/java/com/seriouscompany/business/java/fizzbuzz/packagenamingpackage/interfaces/parameters

The standard layout of JDK projects is the maven layout looks like:

```
src/main/java/         # Application/Library sources
src/main/resources/    # Application/Library resources
src/main/filters/     # Resource filter files
src/main/webapp/      # Web application sources
src/test/java/        # Test sources
src/test/resources/   # Test resources
src/test/filters/    # Test resource filter files
src/it/              # Integration Tests
src/assembly/        # Assembly descriptors
src/site/           # Site documentation
target/             # Generated output (compiled classes, JAR files, etc.)
pom.xml             # Project Object Model (project configuration)
```

You can see it here: https://maven.apache.org/guides/introduction/introduction-to-the-standard-directory-layout.html

The standard maven layout has things I have never seen in a JDK project in the maybe 8 years I have been an engineer, like `src/main/filters`.

The addition of Reverse Domain Convention, having projects named like com.example, makes this directory layout always seem very verbose, but you don't want a conflict on namespaces.

Kotlin actually has a very sane approach to this as well. If you have a package called `org.example.kotlin.network.socket`, and `org.example.kotlin` is common, it should be stored in `network/socket`
ref: https://kotlinlang.org/docs/coding-conventions.html#directory-structure

I want to see if I can create a single directory layout in kotlin and gradle copying leinengen's layout.

So I created a kotlin project with not too much effort with a flat directory layout. I created a `test-resources` folder as leiningen uses profiles to manage that.
ref: https://stackoverflow.com/questions/33482647/leiningen-missing-test-resources

It was surprisingly easy, and I only get one intellij warning, which is slightly annoying. The jar file contains all the metadata that it should have too.

Let's try with bazel...

I got it working with bazel with slightly more effort, mostly because I don't know it that well.

So that it's. I got a simpler jdk layout working with kotlin, and bazel in a few hours.

The project using kotlin, gradle, and bazel here:
https://github.com/jordanst3wart/alternative-jdk-layout

The way to setup gradle to use different directories is just:

```
sourceSets {
    main {
        kotlin {
            srcDirs("src") // use src in the base
        }
        resources {
            srcDirs("resources") // use resources in the base
        }
    }
    test {
        kotlin {
            srcDirs("test") // use test in the base
        }
        resources {
            srcDirs("test-resources")
        }
    }
:
```
