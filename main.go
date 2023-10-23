package main

import (
	"os"

	"github.com/argonautts/golang-blockchain/cli"
)

func main() {
	// Используем defer, чтобы гарантировать выход из программы с кодом 0 при завершении main().
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}

/*
Блокчейн, или "цепочка блоков", - это технология децентрализованного учёта данных, представляющая собой постоянно растущий список записей, называемых блоками, которые связаны и защищены с помощью криптографии. Каждый блок обычно содержит криптографический хэш предыдущего блока, временную метку и данные транзакций.

Основные характеристики и преимущества блокчейна:

1. **Децентрализация**: В традиционной системе, такой как банковская, все данные хранятся на централизованном
сервере. В блокчейне данные распределены по всей сети, и у каждого участника есть копия всего блокчейна.

2. **Прозрачность**: Изменения в публичном блокчейне могут быть просмотрены всеми участниками сети. Это
обеспечивает прозрачность и верифицируемость транзакций.

3. **Неизменность**: После добавления данных в блокчейн их нельзя изменить ретроактивно. Это обеспечивает
целостность и подлинность данных.

4. **Консенсусные алгоритмы**: Для подтверждения и добавления транзакций в блокчейн используются различные
методы достижения консенсуса, такие как Proof of Work (PoW) и Proof of Stake (PoS).

5. **Безопасность**: Благодаря криптографическим алгоритмам и децентрализованной структуре блокчейн является
устойчивым к мошенничеству и атакам.

Наиболее известное применение блокчейна - это криптовалюты, такие как Bitcoin и Ethereum. Однако технология
блокчейна может быть использована в различных сферах: управлении идентификацией, учёте сделок в торговле,
логистике, голосовании и многих других.

Важно отметить, что существуют разные типы блокчейнов: публичные (открытые для всех), частные (доступные
	только для определенных участников) и консорциумные (управляемые группой организаций).
*/
