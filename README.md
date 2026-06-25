<p align="center">
	<img width="80%" src="https://github.com/hexhowells/Quickstart/blob/main/logo.png" alt="Quickstart tool logo">
</p>

Quickstart is a linux command line utility that lets you build man pages containing only the essential commands you need, making it super fast to recall tool usage for only the commands you use.


## Quickstart
First create a new quickstart page

```bash
quickstart new <tool-name>
```

This will create a new template page and open it in your default editor to be customised.

Now we can view the page using
```bash
quickstart <tool-name>
```

We can edit an existing quickstart page with

```bash
quickstart edit <tool-name>
```

If a page is no longer needed, it can be deleted entirely

```bash
quickstart delete <tool-name>
```
