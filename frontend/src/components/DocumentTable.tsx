import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Checkbox } from '@material-ui/core';

interface Props {
  selectedProject: number;
}

const useStyles = makeStyles({
  table: {
    minWidth: 650,
    cursor: 'pointer'
  }
});
export default function DocumentTable({ selectedProject }: Props) {
  const classes = useStyles();

  const documents = ['Document 1', 'Document 2', 'Document 3', 'Document 4', 'Document 5'];

  return (
    <>
      <TableContainer>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell padding="checkbox">
                <Checkbox />
              </TableCell>
              <TableCell>Name</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {documents.map((document, i) => (
              <TableRow key={i} hover={true}>
                <TableCell padding="checkbox">
                  <Checkbox />
                </TableCell>
                <TableCell component="th" scope="row" onClick={() => handleCellClick(i)}>
                  {document}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
