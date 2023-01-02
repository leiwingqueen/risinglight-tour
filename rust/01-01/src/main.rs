use log::debug;
use sqlparser::ast::{Expr, SelectItem, SetExpr, Statement, Value};
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
        match stmt {
            Statement::Query(query) => {
                println!("qqq {:#?}", query.body);
                match query.body {
                    SetExpr::Select(select) => {
                        println!("{:#?}", select);
                        let mut output = String::new();
                        for item in &select.projection {
                            // write!(output, " ").unwrap();
                            match item {
                                SelectItem::UnnamedExpr(Expr::Value(v)) => match v {
                                    // Value::SingleQuotedString(s) => write!(output, "{}", s).unwrap(),
                                    // Value::Number(s, _) => write!(output, "{}", s).unwrap(),
                                    _ => todo!("not supported statement: {:#?}", stmt),
                                },
                                _ => todo!("not supported statement: {:#?}", stmt),
                            }
                            println!("item{:#?}", output);
                        }
                    }
                    _ => todo!("not supported statement: {:#?}", stmt),
                }
            }
            // default handler
            _ => todo!("not supported statement: {:#?}", stmt),
        }
    }
}