use log::debug;
use sqlparser::dialect::PostgreSqlDialect;
use sqlparser::parser::Parser;

fn main() {
    let mut sql = String::new();
    std::io::stdin().read_line(&mut sql).unwrap();
    // use sql parser to transfer the sql to AST
    let stmts = Parser::parse_sql(&PostgreSqlDialect {}, &sql);
    if !stmts.is_ok() {
        return;
    }
    println!("{:#?}", stmts);
    /*while let Ok(ref stmt) = stmts{
        debug!("execute: {:#?}", stmt);
    }*/
    for stmt in stmts.unwrap() {
        debug!("execute: {:#?}", stmt);
    }
}