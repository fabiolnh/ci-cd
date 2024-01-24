# **Guia de Contribuição**

Este é o guia de contribuição do 'ABC'. Você encontrará aqui alguns recursos para ajudá-lo a contribuir.

## **Pré-requisitos**

Confira os requisitos antes de contribuir com o projeto:

### Conventional Commits Git Hook

- Para seguir este tutorial é necessário ter o `npm` instalado.

- Instalar o `git-conventional-commits` de forma global com o comando:

```
npm install --global git-conventional-commits
```

#### Adicionando o git hook de forma global

- Crie uma pasta onde irão ficar armazenados os `hooks` globais

- Configure o `git` para apontar para a pasta com o comando:

```
git config --global core.hooksPath "Endereço da Pasta Criada"
```

- Na pasta criar um arquivo com nome de `commit-msg`

- Colar o seguinte conteúdo e salvar, conforme a [documentação](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13)

```
#!/usr/bin/env sh

commit_message="$1"
# exit with a non zero exit code incase of an invalid commit message

printf "\n"
printf "Verifying conventional commit..."
printf "\n"

# use git-conventional-commits, see https://github.com/qoomon/git-conventional-commits
git-conventional-commits commit-msg-hook "$commit_message"
```

**Obs:** Em sistemas baseados em UNIX é necessário transformar o arquivo em executável usando o comando **chmod +x commit-msg**

# **Como contribuir?**

Consulte as diretrizes para enviar suas alterações.

### **Como você pode contribuir?**

Você pode contribuir com um novo recurso, correção de bugs, documentação e muito mais, confira abaixo:

- **build:** alterações que afetam o sistema de compilação (alterando a configuração do webpack, por exemplo);
- **ci:** alterações em nossos arquivos e scripts de configuração de CI;
- **docs:** referem-se a inclusão ou alteração somente de arquivos de documentação;
- **feat:** tratam adições de novas funcionalidades ou de quaisquer outras novas implantações ao código;
- **fix:** essencialmente definem o tratamento de correções de bugs;
- **perf:** uma alteração de código que melhora o desempenho;
- **refactor:** tipo utilizado em quaisquer mudanças que sejam executados no código, porém não alterem a funcionalidade final da tarefa impactada;
- **style:** alterações referentes a formatações na apresentação do código que não afetam o significado do código, como por exemplo: espaço em branco, formatação, ponto e vírgula ausente etc.);
- **test:** adicionando testes ausentes ou corrigindo testes existentes nos processos de testes automatizados;
- **chore:** atualização de tarefas que não ocasionam alteração no código de produção, mas mudanças de ferramentas, mudanças de configuração e bibliotecas que realmente não entram em produção, `dev dependencies`;
- **merge:** essencialmente merge entre commits
- **revert:** essencialmente reversão para um commit anterior

> **Nota:**
> Se você estiver introduzindo uma `breaking change`, o rodapé da mensagem deve começar com [**`BREAKING CHANGE:`**](https://www.conventionalcommits.org/en/v1.0.0/#commit-message-with-description-and-breaking-change-footer).

### **Forma de escrita do commit**

A mensagem do commit deve ser estruturada da seguinte forma:

```
<tipo>[escopo opcional]: <descrição>

[corpo opcional]

[rodapé(s) opcional(is)]
```

> **Nota:**
> Em relação a descrição, é ideal que o verbo esteja no imperativo (ex: corrige, atualiza, adiciona e etc) e que tudo esteja em minúsculo.
