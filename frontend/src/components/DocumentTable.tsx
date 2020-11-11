import React from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Checkbox } from '@material-ui/core';
import { DocumentDTO } from '../api/apis';

interface Props {
  selectedProject: number;
  documents: DocumentDTO[];
}

const useStyles = makeStyles({
  table: {
    minWidth: 650,
    cursor: 'pointer'
  }
});
export default function DocumentTable({ selectedProject, documents }: Props) {
  const classes = useStyles();

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
                <TableCell component="th" scope="row" onClick={() => handleCellClick(document.id)}>
                  {document.name}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
