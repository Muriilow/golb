;;;
{
    "title": "Hacking 1",
    "description": "Explaining how Im a cool hacker",
    "date": "2025-11-6",
    "author": "Mur1il0w"
}
;;;

# vinaC

VINAc é um programa construído com a linguagem de programação C, que tem como objetivo salvar uma coleção de arquivos comprimidos ou descomprimidos dentro de um arquivo binário.

Este é um software desenvolvido por *Murilo de Paula Bob*, GRR: **20242184**.

### Conteúdo

O projeto possui alguns arquivos fontes, cabeçalhos e um makefile. 

- **main.c**: Execução do programa principal, lê os argumentos mandados pelo usuário e realiza as operações. 
- **vinaC.c - vinaC.h**: Declarações e implementações de cada função responsável por editar e configurar o *archive*.
- **fileSafety.c - fileSafety.h**: Declarações e implementações de funções que realizam operações como leitura ou escrita de arquivo, checando por possíveis erros. 
- **lz.c - lz.h**: Declarações e implementações feitas por *Marcus Geelnard* que realiza as aritméticas necessária para comprimir e descomprimir arquivos. 

1. um 
2. dois 
    - salve 
3. tres
    - salve
        - opa

### Descrição das escolhas de implementação

O programa não utiliza nenhum tipo abstrato de dado para modificar ou salvar informações, apenas movimentações de dados em um arquivo binário foram utilizadas para a formação do programa, como *fwrite()* ou *fread()*. Utilizei o método camelCase para assinar variavéis e PascalCase para assinar as funções e structs.


#### Porque dessa escolha?

Foi o que pareceu mais natural no inicio do desenvolvimento, logo depois se provou uma forma um pouco mais difícil mas não chegou a ser limitante, portanto continuei da mesma forma.


A forma como foram contruidas as funções do VinaC provavelmente seguem a mesma maneira de outros trabalhos, [tenho](https://www.google.com) como função essencial *MoveData()* que movimenta um bloco de dados para uma certa posição no binário.
``` C 
void MoveData(long int start, long int size, long int pos, FILE* binary)
{
    char* buffer = SafeMalloc(size);

    fseek(binary, start, SEEK_SET);

    SafeReadFile(buffer, 1, size, binary);

    fseek(binary, pos, SEEK_SET);
    SafeWriteFile(buffer, 1, size, binary);
    free(buffer);
}
```
A função *MoveFixedData()*, abstrai a complexidade de mover dados em partes.

> ### Por que?
>> Como definição o temos o tamanho máximo do buffer de escrita limitado por um valor. 

``` C
void MoveFixedData(long int start, long int size, long int pos, long int maxSize, FILE* binary)
{
    int end, sizeDiff;

    if(start > pos)
    {
        while(size > maxSize)
        {
            MoveData(start, maxSize, pos, binary);
            start += maxSize;
            pos += maxSize;
            size -= maxSize;
        }
        MoveData(start, size, pos, binary);
        return;
    }

    end = start + size;
    sizeDiff = pos - start;

    while(size > maxSize)
    {
        pos = end - maxSize;
        end -= maxSize;
        MoveData(pos, maxSize, pos + sizeDiff, binary);
        size -= maxSize;
    }
    pos = end - size;
    MoveData(pos, size, pos + sizeDiff, binary);
};
```

> #### Alerta
> No meu programa eu "abstraio" funções essenciais como *malloc()* ou *fwrite()*, encapsulando elas em funções externas que checam se essas operações foram realizadas com sucesso, dessa forma eu não preciso fazer um *if* toda vez que checo se uma operação de *malloc* foi realizada com sucesso. 


Um exemplo de escrita segura:
``` C
void SafeWriteFile(void *buffer, size_t size, size_t count, FILE *file)
{
    if (buffer == NULL || file == NULL)
    {
        fprintf(stderr, "Error: Null pointer passed to SafeWriteFile.\n");
        exit(EXIT_FAILURE);
    }

    /* Perform the write */
    size_t elementsWritten = fwrite(buffer, size, count, file);

    /* Check for errors or incomplete writes */
    if (elementsWritten != count)
    {
        if (ferror(file))
            perror("Error during fwrite");

        exit(EXIT_FAILURE);
    }
}

``` 

Dessa forma meu programa esta constantemente checando possíveis erros, olhando o retorno dessas funções elementares. 

### Problemas enfrentados

Em geral a complexidade de modificar strings e dados binários foi a principal e única dificuldade, principalmente da forma como desenvolvi esses problemas, no inicio entender como realizar o projeto era bem abstrato, tive dificuldades para dar os primeiros passos, mas depois que entendi bem funções como *fseek(), ftell(), strncpy(), strcmp()...*, como interpretar entradas do usuário por meio de argv[] e argc, e como o sistema operacional possui um "ponteiro" para arquivos, a resolução dos problemas ficaram simples. 

#### Possíveis bugs 

Caso o usuário utilizar o programa corretamente nenhum erro é esperado, ou seja:

``
./vina -i archive.vc member1 member2 ...
``

Mas caso o usuário passe um arquivo normal **já existente** como o *archive.vc* a resposta do programa será indefinida , isto se deve porque ao criar um *archive.vc* o programa escreve uma struct diretório contendo a quantidade de membros presentes no *archive.vc*, caso o usuário passe um arquivo comum como um *archive.vc* o programa tentará ler um diretório, mas o conteúdo do diretório ficará indefinido, seria possivel fazer uma função que checasse se o arquivo passado como *archive.vc* é valido, mas devido ao tempo e a presunção que não será levado em consideração a falha humana na hora do correção optei por ignorar.  

Outro caso interessante de comentar, mas que não é um bug e sim uma implementação ruim, é quando queremos modificar um membro pequeno, em um *archive.vc* muito grande, a forma como implementei utiliza o tamanho do membro a ser modificado como limitador do buffer, caso eu queria mover um membro de 100 bytes em um binário de 100000 bytes, a quantidade de *mallocs* será grande levando a um programa mais lento.

Seria possível resolver este problema armazenando o tamanho do maior membro na struct diretório, e utilizando sempre esse valor para o buffer, infelizmente so percebi essa má implementação na fase final de teste perto da entrega (momentos antes de escrever o README). 

### Considerações finais 

Meu programa performa suas operações de forma segura, sem erros de alocação dinâmica e fazendo o tratamento de bugs da melhor maneira possível, não foi utilizada nenhuma lista ou array de structs para performar as operações necessárias e o ritmo de desenvolvimento seguiu uma ordem bem natural, algumas implementações poderiam ser mais robusta e rapidas mas que devido ao tempo, e as regras do enunciado não foi possível deixar da melhor maneira. 


### This is a heading example 

```C 
function void snippet(int day)
{
    printf("Lovely day huh");
    return;
}
``` 
