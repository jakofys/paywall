# ![Logo](./assets/Paywall.png)

Paywall is a CLI tool for basic operation calculation based on Lightning Network protocol write in go.

---

## Requirements

- Golang 1.13 or more
- Docker 20.10.7

## Installation

> This default installation allow you to run locally the application.

Make in first step following command in project folder:

```bash
docker build Dockerfile && docker run bitcoin-regtest
```

If the output show you that ok, then you can make:

```bash
go get
```

## Build

To build the CLI tool, you have two possibilities.

**First** and default behavior, run the `go build cmd/paywall/paywall.go paywall` script and then, build basicly your program according to current OS.

**Second** way is to make you own build recipe executing:

```bash
$ GOOS=[os_target] GOARCH=[cpu_arch] go build paywall
```

Which `os_target` is OS platform on you want to build an executable, and `cpu_arch` is the architecture of the instruction set according to CPU.

You can refer available env argument for **GOOS** and **GOARCH** on [this repository](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)

Finally, you can execute you distributable application named `paywall` previously build using the command

```bash
./paywall
```

## Running

For run the CLI tool, just make:

```bash
go run /cmd/paywall/paywall.go
```

---

## Use

Paywall CLI tool allow you to make basic calcul like **substraction** and **add up** with two member.
The format of CLI command is:

```bash
paywall [command] [member_one] [member_two] [--mainnet]
```

> If you've not build tool, replace `paywall` keyword example to `go run /cmd/paywall/paywall.go`

<img src="https://img.icons8.com/emoji/452/exclamation-mark-emoji.png" width=16> Each member must be a integer.

Then  calculation operate successfully, a transaction amount 1msat is credit to the user and generate a bill.

When transaction operate successfully, the result is getting to the user. If not, an error message is displayed to the user explain him to specify a creditable crypto-currency wallet.

### Option

The only option is called `--mainnet` to determine if the computation will execute on the **mainnet**. By default, it executing on `testnet` environment.

### Substract

```bash
paywall substract [member_one] [member_two]
```

example:

```bash
paywall substract 10 5
# Output 10-5=5
```

### Add

```bash
paywall add [member_one] [member_two]
```

example:

```bash
paywall add 10 5
# Output 10+5=15
```

## Ressoures

[Cobra Go](https://github.com/spf13/cobre): CLI implementation tool write in go

[Viper Go](https://github.com/spf13/viper): Configuration package write in go

[Lightning Network](https://lightning.network/): Is a way to make instant micro-transaction at low cost impact with smart contract in crypto currency context or get on this [video](https://www.youtube.com/watch?v=rrr_zPmEiME) of [Symply Explained](https://www.youtube.com/channel/UCnxrdFPXJMeHru_b4Q_vTPQ) channel.
