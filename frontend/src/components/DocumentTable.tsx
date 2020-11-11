import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';

import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Checkbox, TableSortLabel } from '@material-ui/core';
import { DocumentDTO, getDocumentByDocumentId } from '../api/apis';

interface Props {
  documents: DocumentDTO[];
}

const useStyles = makeStyles({
  table: {
    minWidth: 650,
    cursor: 'pointer'
  }
});

const sortByName = (documents: DocumentDTO[], isDescending: boolean) => {
  function getComperator(a: string, b: string) {
    if (a > b) {
      return 1;
    }
    if (a < b) {
      return -1;
    }
    return 0;
  }

  return documents.sort((a, b) => {
    if (isDescending) {
      return getComperator(a.name, b.name);
    }
    return -getComperator(a.name, b.name);
  });
};

const handleCellClick = async (doc: DocumentDTO) => {
  const myDocument = await getDocumentByDocumentId(doc.id);
  const blopDocument = new Blob([myDocument], { type: 'octet/stream' });
  const documentUrl = window.URL.createObjectURL(blopDocument);

  const link = document.createElement('a');
  link.href = documentUrl;
  link.setAttribute('download', doc.name);
  document.body.appendChild(link);
  link.click();
};

export default function DocumentTable({ documents }: Props) {
  const classes = useStyles();
  const [isDescending, setIsDescending] = useState(false);
  const [sortedDocuments, setSortedDocuments] = useState<DocumentDTO[]>([]);

  useEffect(() => {
    setSortedDocuments(sortByName(documents, isDescending));
  }, [documents, isDescending]);

  return (
    <>
      <TableContainer>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell padding="checkbox">
                <Checkbox />
              </TableCell>
              <TableCell>
                <TableSortLabel active direction={isDescending ? 'desc' : 'asc'} onClick={() => setIsDescending(current => !current)}>
                  Name
                </TableSortLabel>
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {sortedDocuments.map((document, i) => (
              <TableRow key={i} hover={true}>
                <TableCell padding="checkbox">
                  <Checkbox />
                </TableCell>
                <TableCell component="th" scope="row" onClick={() => handleCellClick(document)}>
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
